/**************************************************************
This examples will help you understand how do you declare
a nested struct in another structure.

There are two format for it.
1. Nested a single structure
2. Nested a array of structure

Please read it carefully to understand it
**************************************************************/
// Golang program to illustrate
// the nested structure
package main

import "fmt"

// Creating structure
type Student struct {
	name   string
	branch string
	year   int
}

// Creating nested structure
type Tutor struct {
	name    string
	subject string
	exp     int
	details Student
}

// Creating a array of nested structure
type Teacher struct {
	name    string
	subject string
	exp     int
	details []Student
}

func main() {

	// Initializing the fields
	// of the structure
	result := Teacher{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		details: []Student{
			{
				name:   "Bongo",
				branch: "CSE",
				year:   2,
			},
			{
				name:   "Hua Van Thong",
				branch: "Uni",
				year:   6,
			},
		},
	}

	// Display the values
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", result.name)
	fmt.Println("Subject: ", result.subject)
	fmt.Println("Experience: ", result.exp)

	fmt.Println("\nDetails of Students for Teacher")
	fmt.Println("Student's name: ", result.details[0].name)
	fmt.Println("Student's branch name: ", result.details[0].branch)
	fmt.Println("Year: ", result.details[0].year)

	fmt.Println("Student's name: ", result.details[1].name)
	fmt.Println("Student's branch name: ", result.details[1].branch)
	fmt.Println("Year: ", result.details[1].year)

	tutor := Tutor{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		details: Student{
			name:   "Bongo",
			branch: "CSE",
			year:   2,
		},
	}

	// Display the values
	fmt.Println("#########################################")
	fmt.Println("Details of the Tutor")
	fmt.Println("Teacher's name: ", tutor.name)
	fmt.Println("Subject: ", tutor.subject)
	fmt.Println("Experience: ", tutor.exp)

	fmt.Println("\nDetails a Student for tutor")
	fmt.Println("Student's name: ", tutor.details.name)
	fmt.Println("Student's branch name: ", tutor.details.branch)
	fmt.Println("Year: ", tutor.details.year)
}
