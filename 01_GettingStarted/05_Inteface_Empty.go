// Reference: https://go.dev/tour/methods/14
package main

import "fmt"

/*
*****************************************************************************
Define a data-type is a structure
*****************************************************************************
*/
type Description struct {
	ID                int
	Name              string
	DescriptionDetail string
}

/*
*****************************************************************************
desribe() function have an argunment as a empty interface
*****************************************************************************
*/
func describe(i ...interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
*****************************************************************************
main
*****************************************************************************
*/
func main() {

	// Method 1: We can pass a empty
	var i interface{}
	describe(i)

	// Method 2: We can pass any data type to describe() func
	i = 42
	describe(i)

	i = "hello"
	describe(i)

	d := Description{
		ID:                1,
		Name:              "hvthong",
		DescriptionDetail: "Try to test Interface with struct data-type",
	}
	describe(d)

	// Method 3: Update function to insert multiple arg
	describe(d, i)
}
