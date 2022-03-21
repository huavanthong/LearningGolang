// Reference:
// https://stackoverflow.com/questions/59976812/empty-interfaces-in-golang
package main

import (
	"fmt"
)

/*
*****************************************************************************
Define method inside a inteface provide all sub-structure
*****************************************************************************
*/
type NamedType interface {
	getName() string
}

/*
*****************************************************************************
Implmenent Dog structure and methods for this Dog struct
*****************************************************************************
*/
type Dog struct {
	name string
}

func (d Dog) getName() string {
	return d.name
}

/*
*****************************************************************************
Implmenent Cat structure and methods for this Cat struct
*****************************************************************************
*/
type Cat struct {
	name string
}

func (c Cat) getName() string {
	return c.name
}

/*
*****************************************************************************
Implement function to getName() from both Dog and Cat structure.
However, you can see that:
	+ sayHiIssue() function can't getName
	+ To fix it, we implement sayHi() function
*****************************************************************************
*/
func sayHiIssue(i interface{}) {

	fmt.Println(i, "says hello")

	// Not understanding this error message
	fmt.Println(i.name) //  i.name undefined (type interface {} is interface with no methods)
}

// This function will print says Hello to all structures that implement interface NameType
// The meaning is as same as parent contain all methods from children
func sayHi(i NamedType) {

	fmt.Println(i, "says hello")

	// To fix the above issue
	// Step 1: define a inteface to getName() for both struct Dog and Cat
	// Step 2: Each struct, we declare a method to getName() from it.
	fmt.Println(i.getName())
}

/*
*****************************************************************************
main
*****************************************************************************
*/
func main() {
	d := Dog{"Sparky"}
	c := Cat{"Garfield"}

	sayHi(d) // {Sparky} says hello
	sayHi(c) // {Garfield} says hello
}
