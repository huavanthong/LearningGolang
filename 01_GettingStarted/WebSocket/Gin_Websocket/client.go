package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// Địa chỉ của WebSocket server
	addr := "ws://localhost:8080/ws"

	// Thực hiện kết nối đến WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Gửi thông điệp đến server
	message := []byte("Hello, server!")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal(err)
	}

	// Đọc phản hồi từ server
	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Received: %s", response)
}
