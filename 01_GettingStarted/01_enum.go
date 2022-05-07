// Reference: https://techmaster.vn/posts/35010/huong-dan-co-ban-ve-enums-trong-go
package main

import "fmt"

// Step 1: declare a type
type Weekday int

// Step 2: declare a list of value const for your type
const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

/***************************************************************/
// Step 3: create some behaviors for enum
/***************************************************************/
// String(): convert day to string
func (day Weekday) String() string {
	// khai báo một mảng các string
	// toán tử ... để đếm số phần tử
	// số phần tử của mảng là (7)
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}

	// `day`: là một trong các giá trị của hằng số Weekday.
	// Nếu hằng số là Sunday, thì day có giá trị là 0.
	// Bắt lỗi trong trường hợp `day` nằm ngoài khoảng của Weekday
	if day < Sunday || day > Saturday {
		return "Unknown"
	}
	// trả về tên của 1 hằng số Weekday từ mảng names bên trên
	return names[day]
}

// Weekend(): determine that day in range weekend
func (day Weekday) Weekend() bool {
	switch day {
	// Nếu day là ngày cuối tuần:
	case Sunday, Saturday:
		return true
	// Nếu day là ngày trong tuần:
	default:
		return false
	}
}

func main() {

	var monday Weekday = Monday
	fmt.Println(monday.String())
	fmt.Println(monday.Weekend())

}
