# Introduction
WebSocket là một giao thức được sử dụng để thiết lập một kết nối liên tục (persistent connection) giữa máy khách (client) và máy chủ (server) để truyền tải dữ liệu hai chiều (full-duplex communication). Thường thì nó được sử dụng để tạo ra các ứng dụng truyền tải dữ liệu realtime, như chat, game, giao dịch trực tuyến và các ứng dụng web tương tác. Với WebSocket, các thông tin được truyền tải liên tục và theo thời gian thực, cho phép các ứng dụng web thực hiện các tác vụ như gửi thông báo, cập nhật dữ liệu và tương tác người dùng một cách nhanh chóng và hiệu quả hơn.

# Websocket in Golang
Trong Golang, có nhiều thư viện hỗ trợ việc sử dụng Websocket để xây dựng ứng dụng livestream. Một số thư viện phổ biến như sau:

1. Gorilla Websocket: Đây là một thư viện websocket phổ biến được sử dụng rộng rãi trong Go. Thư viện này hỗ trợ cả client và server-side WebSocket.

2. nhooyr.io/websocket: Thư viện này hỗ trợ việc xây dựng cả client và server-side WebSocket. Nó cho phép xác thực, mã hóa và nén dữ liệu trước khi truyền qua websocket.

3. Gobwas/ws: Thư viện này hỗ trợ việc xây dựng server-side WebSocket. Nó rất nhẹ và nhanh, được tối ưu hóa để xử lý các kết nối lớn.

