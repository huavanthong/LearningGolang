package main

import "fmt"

type data int

func main() {
	fmt.Println("Run case 1:")
	var i interface{}

	i = 42
	value1, ok := i.(int)
	if ok {
		fmt.Println("The value is an int:", value1)
	} else {
		fmt.Println("The value is not an int")
	}

	fmt.Println("Run case 2:")
	var test interface{}

	test = data(42)

	value, ok := test.(data)
	if ok {
		fmt.Println("The value is an data:", value)
	} else {
		fmt.Println("The value is not an int")
	}

}
