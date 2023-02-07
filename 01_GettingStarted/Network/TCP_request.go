package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return
	}
	defer conn.Close()

	// Send data to the server
	request := "Hello, World!\n"
	fmt.Print("Sending: ", request)
	fmt.Fprintf(conn, request)

	// Read the response from the server
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// Print the received response
	fmt.Print("Received: ", response)
}
