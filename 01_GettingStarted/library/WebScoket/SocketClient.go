package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the remote TCP server
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	// Send data to the server
	message := "Hello, server"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err.Error())
		os.Exit(1)
	}

	// Receive data from the server
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error receiving data:", err.Error())
		os.Exit(1)
	}

	// Print the received data
	fmt.Println("Received data:", string(buf[:n]))

	// Close the connection
	conn.Close()
}
