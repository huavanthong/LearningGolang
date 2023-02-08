package main

import "fmt"

/*
In Go, make and new are two different built-in functions that are used to allocate memory and initialize values.
Here is the difference between the two:
    * make: make is used to create slices, maps, and channels. It takes a type and the number of elements as arguments
			and returns a pointer to a newly created, empty data structure of the specified type.
			make also initializes the underlying data structure,
			such as allocating memory for the map and setting up the hash table.

	* new: new is used to allocate memory for a new value of a specified type.
			It returns a pointer to the newly allocated value, which is set to zero.
			Unlike make, new does not initialize the underlying data structure.

In general, make is used for creating and initializing Go's built-in data structures, such as slices, maps, and channels. new is used for allocating memory for custom types and values.

Here is an example of using make to create a slice:
	| s := make([]int, 5)

And here is an example of using new to allocate memory for a custom type:
	| p := new(Person)
*/
type Person struct {
	ID   int
	Name string
}

func main() {
	fmt.Printf("Example for make\n")
	// Create a map with initial capacity of 10
	m := make(map[string]int, 10)

	// Add some key-value pairs to the map
	m["first"] = 1
	m["second"] = 2

	// The map has been initialized and is ready to use
	fmt.Println(m)
	// Output: map[first:1 second:2]

	var p Person
	ptr := new(Person)
	p.ID = 1
	p.Name = "Van Thong"

	fmt.Println("ID: %d, Name: %s", p.ID, p.Name)
	fmt.Println("Before point a variable")
	fmt.Println("ID: %d, Name: %s", ptr.ID, ptr.Name)
	fmt.Println("After point a variable")
	ptr = &p
	fmt.Println("ID: %d, Name: %s", ptr.ID, ptr.Name)

}
