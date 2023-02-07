package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// Create a listener for incoming connections
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on 0.0.0.0:8080")

	// Listen for incoming connections indefinitely
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}

		// Handle the incoming connection in a new goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	// Read data from the connection
	request, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// Print the received data
	fmt.Print("Received: ", request)

	// Write the response back to the connection
	response := strings.ToUpper(request)
	conn.Write([]byte(response))
}
