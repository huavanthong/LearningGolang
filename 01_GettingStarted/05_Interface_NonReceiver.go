package main

import (
	"errors"
	"fmt"
	"strings"
)

/*************************************************************************************************/
/* Parent structure using interface
/*************************************************************************************************/
// Go Interface - `Shape`
type Shape interface {
	Uppercase(string) (string, error)
	Area() float64
	Perimeter() float64
}

/*************************************************************************************************/
/* Rectangle will be the Shape'children if it implements all Shape's methods
/*************************************************************************************************/
// Struct type `Rectangle` - implements the `Shape` interface by implementing all its methods.
type Rectangle struct {
	name          string
	Length, Width float64
}

/********* Global variable *********/
var ErrEmpty = errors.New("empty string")

/********* Methods inherit *********/
// Methods inherit from Shape
func (Rectangle) Uppercase(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

/********* Methods utility *********/
// Set name this rectangle by using pointer
func (r *Rectangle) setName(name string) {
	(*r).name = name
}

// Set width by using pointer
func (r *Rectangle) setWidth(width float64) {
	(*r).Width = width
}

// Set length by using pointer
func (r *Rectangle) setLength(length float64) {
	(*r).Length = length
}

// show info by using value
func (r Rectangle) show() {
	fmt.Println("Length : ", r.Length)
	fmt.Println("Width  : ", r.Width)
}

func main() {

	var rect Rectangle
	rect.setName("hinhvuong")
	rect.setLength(10)
	rect.setWidth(12)
	fmt.Printf("Rect Type = %T, Rhape Value = %v\n", rect, rect)
	fmt.Printf("Name: %s, Area = %f, Perimeter = %f\n", rect.Uppercase(rect.name), rect.Area(), rect.Perimeter())
}
