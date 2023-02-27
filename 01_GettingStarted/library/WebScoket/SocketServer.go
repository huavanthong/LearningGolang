package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP connection to WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		// Handle WebSocket messages
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}
			log.Printf("Received message: %s\n", message)

			// Send response
			response := fmt.Sprintf("Received message: %s", message)
			err = conn.WriteMessage(messageType, []byte(response))
			if err != nil {
				log.Println(err)
				break
			}
		}
	})

	// Start WebSocket server
	log.Println("Starting WebSocket server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
