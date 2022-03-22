package main

import "fmt"

type Student struct {
	Id      int    // Student'Id
	Name    string // Student'name
	Subject string // Subject learning
}

func main() {

	fmt.Println("==== Method 1: Create structure variable ====")
	var test Student
	test.Id = 1
	test.Name = "hvthong"
	test.Subject = "math"
	fmt.Printf("Struct Variable: Check Id: %d\n", test.Id)
	fmt.Printf("Struct Variable: Check Name: %s\n", test.Name)
	fmt.Printf("Struct Variable: Check Subject: %s\n", test.Subject)
	fmt.Printf("Struct Variable: Check address: %p\n", &test)

	fmt.Println("==== Method 2: Create structure pointer ====")
	var ptr_test *Student = &test
	fmt.Printf("Struct Pointer: Check value %d\n", test.Id)
	fmt.Printf("Struct Pointer: Check address %p\n", ptr_test)

	fmt.Println("==== Method 3: Initialize a structure ====")
	student := Student{}
	fmt.Printf("Struct Pointer: Check value %d\n", student.Id)
	fmt.Printf("Struct Pointer: Check address %p\n", &student)

	fmt.Println("==== Method 2: Create structure pointer ====")
	var ptr_student *Student = &student
	fmt.Printf("Struct Pointer: Check value %d\n", ptr_student.Id)
	fmt.Printf("Struct Pointer: Check address %p\n", ptr_student)

	fmt.Println("==== Method 4: Create pointer point a new memory area ====")
	p := new(Student)
	fmt.Printf("Struct new: Check value %d\n", p.Id)
	fmt.Printf("Struct new: Check address %p\n", p)

	fmt.Println("==== Method 5: Consider using allocation with new as a pointer ====")
	p = &test
	fmt.Printf("Struct Variable: Check Id: %d\n", p.Id)
	fmt.Printf("Struct Variable: Check Name: %s\n", p.Name)
	fmt.Printf("Struct Variable: Check Subject: %s\n", p.Subject)
	fmt.Printf("Struct Variable: Check address: %p\n", p)
}
