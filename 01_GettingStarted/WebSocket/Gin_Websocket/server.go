package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	r := gin.Default()

	// Khởi tạo WebSocket
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	r.GET("/ws", func(c *gin.Context) {
		// Thực hiện lấy kết nối WebSocket từ HTTP
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		// Lặp vô hạn để đọc thông tin từ WebSocket
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("Received: %s", msg)

			// Gửi thông tin trả lời về WebSocket
			err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
			if err != nil {
				log.Println(err)
				return
			}
		}
	})

	r.Run(":8080")
}
