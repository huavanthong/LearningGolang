package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

/*
Error: read tcp [::1]:53548->[::1]:8080: wsarecv: An established connection was aborted by the software in your host machine.

This error message indicates that a network connection has been terminated unexpectedly. There could be several reasons for this error, including:

    1. The server may have closed the connection.
    2. There may be a network issue or the connection may have timed out.
    3. The host machine may have closed the connection.

Root cause:
	1. Beacuse your firewall detect the maintain connection on 8080 port, so it wll disconnect that connection.

Solution:

If the solution I provided did not resolve the issue, here are a few other steps that you can try:
    1. Check the logs on the server side to see if it is closing the connection or if there are any other issues on the server side.
	2. If the server is closing the connection, you may need to modify the server-side code to keep the connection open.
	3. If there is a network issue, you may need to troubleshoot the network or check for any firewall rules that could be blocking the connection.
	4. If the host machine is closing the connection, you may need to check for any antivirus software or firewall rules that could be causing the issue.
	5. Try to reproduce the issue with a smaller, simplified example to isolate the problem.
	6. Consider using a different networking library or a different networking protocol (e.g. WebSockets instead of TCP) if the issue persists.


*/
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading", err.Error())

			//handle error
			conn.SetDeadline(time.Now().Add(time.Duration(5) * time.Second))

			return
		}
		fmt.Print("Message from server: " + message)
	}
}
