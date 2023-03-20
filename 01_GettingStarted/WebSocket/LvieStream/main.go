package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Tạo một broadcast type để quản lý các clients và messages.
type broadcast struct {
	// Một đối tượng để gửi message đến các clients khác nhau.
	clients map[*client]bool

	// Một kênh để gửi message đến tất cả các clients.
	broadcast chan []byte

	// Một đối tượng để đồng bộ hóa truy cập đến danh sách clients.
	mutex sync.RWMutex
}

// Tạo một client type để lưu trữ thông tin về một client kết nối.
type client struct {
	// Một đối tượng websocket cho kết nối.
	conn *websocket.Conn

	// Một kênh để gửi message đến client.
	send chan []byte

	// Một đối tượng broadcast mà client đang kết nối tới.
	b *broadcast
}

// Hàm chạy cho mỗi client để đọc và gửi message.
func (c *client) run() {
	defer func() {
		c.b.mutex.Lock()
		delete(c.b.clients, c)
		close(c.send)
		c.b.mutex.Unlock()
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		// Gửi message đến kênh broadcast.
		c.b.broadcast <- message
	}
}

// Hàm khởi tạo một broadcast và bắt đầu lắng nghe các message đến và đi.
func newBroadcast() *broadcast {
	return &broadcast{
		clients:   make(map[*client]bool),
		broadcast: make(chan []byte),
	}
}

func (b *broadcast) run() {
	for {
		// Đợi message mới từ kênh broadcast.
		message := <-b.broadcast

		// Duyệt qua tất cả các clients đang kết nối và gửi message đến họ.
		b.mutex.RLock()
		for client := range b.clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(b.clients, client)
			}
		}
		b.mutex.RUnlock()
	}
}

// Hàm chạy khi có một kết nối mới đến server.
func serveWs(b *broadcast, w http.ResponseWriter, r *http.Request) {
	// Thực hiện một upgrade của HTTP connection sang một WebSocket connection.
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		log.Println(err)
		return
	}

	// Tạo một client mới và thêm vào danh sách clients.
	client := &client{
		conn: conn,
		send: make(chan []byte, 256),
		b:    b,
	}
	b.mutex.Lock()
	b.clients[client] = true
	b.mutex.Unlock()

	// Bắt đầu chạy goroutine để đọc và gửi message.
	go client.run()
}

func main() {
	// Tạo một broadcast để quản lý các clients và messages.
	b := newBroadcast()

	// Bắt đầu goroutine để lắng nghe các message mới.
	go b.run()

	// Cài đặt route để handle các WebSocket requests.
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(b, w, r)
	})

	// Start server.
	log.Println("Server starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
