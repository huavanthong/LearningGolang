package main

// Điều đầu tiền khi làm việc với Golang đó chính là import các library trong Golang.
// Qua đó chúng ta phải hiểu ý nghĩa của việc import các library vào qua tutorial
// * #Names: https://www.meisternote.com/app/note/lqIBK1FcK8-g/names
// Ý nghĩa của việc import các library là ta có thể sử dụng các methods trong library đó bằng chính names ta import vào.
// Xem ví dụ trong main() function.
import (
	"fmt"
	"net/http"
)

// Ta xây dựng hàm Handler với các input parameter là Response và Request.
// Để hiểu process của một Web Application ta có thể refer:
//* #HTTP: https://www.meisternote.com/app/note/ofKerE93kWmY/http-overview
func handler(w http.ResponseWriter, request *http.Request) {
	// Bằng việc áp dụng Fprintf() tương tự như C programming. Ta có thể nhập nhiều input để write vào Response Writer mà ta sẽ gửi lên client.
	// Lưu ý ở Request thì sẽ chứa tất cả các info cần thiết để handle request đó.
	// Ta có thể sự interchangeablly trong URN, URL và URI
	// Refer: https://www.meisternote.com/app/note/TaMdVSh8irub/1-7-uri
	fmt.Fprintf(w, "Hello world", request.URL.Path[1:])
}

// Main function.
func main() {

	//Create Handle function để handle URI đó.
	http.HandleFunc("/", handler)
	// Create listen port 80808
	http.ListenAndServe(":8080", nil)
}
