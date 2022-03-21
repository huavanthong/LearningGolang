// Reference: https://www.geeksforgeeks.org/how-to-access-interface-fields-in-golang/
// Golang program to access the interface fields
package main

import (
	"fmt"
	"strings"
)

// Declare course structure
type Course struct {
	name string
}

// Declare contest structure
type Contest struct {
	name string
}

// Declare interface
type Name interface {
	getName() string
}

// getName function for course
func (a Course) getName() string {

	return a.name

}

// getName function for contest
func (b Contest) getName() string {

	return b.name

}

// Compare course and contest name.
// Name is interface type
func name_compare(course, contest Name) bool {

	return contest.getName() == course.getName()

}

/*
*****************************************************************************
We will use Variadic method to implement a function join string
*****************************************************************************
*/
func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

/*
*****************************************************************************
main
*****************************************************************************
*/
func main() {

	var course_name, contest_name string

	// Get the course name from user
	fmt.Println("Enter course name: ")
	fmt.Scan(&course_name)

	// Get the contest's name from user
	fmt.Println("Enter contest name: ")
	fmt.Scan(&contest_name)

	// Create structure of course
	course := Course{course_name}

	// Create structure of contest
	contest := Contest{contest_name}

	// Call interface function to compare names
	fmt.Print("Is same subjects in course and contest: ", name_compare(course, contest), "\n")

	fmt.Println("Joined name: " + joinstr(course.getName(), contest.getName()))

}
