package main

import "fmt"

func main() {

	fmt.Println("\n==== Case 1: Use maps to store key-value ====\n")
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	// Assigning and fetching map values looks syntactically just like doing the same for arrays
	// and slices except that the index doesn't need to be an integer.

	offset := timeZone["EST"]
	fmt.Printf("offset EST: %d\n", offset)
	fmt.Print(timeZone)

	fmt.Println("\n==== Case 2: Initialize a maps using make ====\n")
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	fmt.Println("\n==== Case 3: Use maps to check a exist value in maps ====\n")
	// A set can be implemented as a map with value type bool.
	// Set the map entry to true to put the value in the set, and then test it by simple indexing.
	attended := map[string]bool{
		"Ann":     true,
		"Joe":     true,
		"hvthong": true,
		"nam":     false,
	}
	person := "giaolinh"

	if attended[person] { // will be false if person is not in the map
		fmt.Println(person, "was at the meeting")
	}
	fmt.Printf("Because %s is not map:\n", person)
	fmt.Print(attended)

	fmt.Println("\n==== Case 4: To delete a map entry ====\n")
	delete(timeZone, "PDT") // Now on Standard Time
	fmt.Print(timeZone)

}
