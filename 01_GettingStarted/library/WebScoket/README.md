# ChatGPT Introduction
WebSocket là một giao thức truyền tải dữ liệu hai chiều (full-duplex) trên một kết nối duy nhất giữa máy khách và máy chủ, cho phép hai bên có thể giao tiếp với nhau một cách thời gian thực. WebSocket được sử dụng trong các ứng dụng yêu cầu giao tiếp trực tiếp, ví dụ như trò chuyện trực tuyến, trò chơi trực tuyến, cập nhật trực tiếp, …

Để sử dụng WebSocket trong HTTP, ta cần sử dụng một WebSocket server để thực hiện kết nối giữa client và server. Một số thư viện WebSocket phổ biến như:

    * Spring Websocket: cung cấp cơ chế xử lý các kết nối WebSocket trong ứng dụng Spring.
    * Socket.IO: cung cấp cơ chế xử lý các kết nối WebSocket trong ứng dụng Node.js.

Để thực hiện kết nối đến một WebSocket server, ta cần sử dụng một thư viện WebSocket trên phía client. Các thư viện WebSocket phổ biến bao gồm:

    * WebSocket API: một API được tích hợp sẵn trong các trình duyệt Web để hỗ trợ kết nối WebSocket trên phía client.
    * SockJS: cung cấp một API đơn giản để kết nối đến một WebSocket server, hỗ trợ các trình duyệt không hỗ trợ WebSocket.

# Ideas
Để xây dựng được Websocket này, chúng ta cần phải hiểu như sau:
1. Đầu tiên, xây dựng một máy chủ WebSocket, nơi tạo ra endpoint và thực hiện các nghiệp vụ business trên nó. Ta có thể xây dựng Websocket bằng bất kì ngôn ngữ nào chẳng hạn như Golang, Java, NodeJS...
    * Trong đó, ta xây dựng một business như sau:
        * 
2. Sau đó, ta xây dựng một máy khách Websocket,chúng ta có thể sử dụng Golang, Javascript hoặc bất kỳ ngôn ngữ nào hỗ trợ WebSocket. 
Quan trọng là chúng ta cần phải xây dựng business như thế nào trên Websocket của chúng ta.

# Getting Started
**Step 1:** Tạo một máy chủ WebSocket bằng Golang
```golang
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

```
Trong đó:  
* Một biến upgrader kiểu websocket.Upgrader để cấu hình kích thước bộ đệm cho việc đọc và ghi các thông điệp WebSocket. Sau đó, chúng ta định nghĩa một hàm xử lý các kết nối WebSocket mới.
* Trong hàm xử lý này, chúng ta sử dụng upgrader để nâng cấp kết nối HTTP ban đầu sang một kết nối WebSocket. Sau đó, chúng ta đọc các thông điệp WebSocket từ máy khách, hiển thị chúng trong log và gửi lại các thông điệp này cho máy khách.
* Cuối cùng, chúng ta bắt đầu máy chủ WebSocket bằng cách gọi hàm http.ListenAndServe() và chạy máy chủ trên cổng 8080.

**Step 2:** Tạo một kết nối WebSocket từ máy khách.
```golang
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

```