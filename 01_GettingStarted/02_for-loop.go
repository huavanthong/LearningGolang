// Reference: https://www.callicoder.com/golang-control-flow/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Declare a array
	var arr [5]int

	fmt.Println("===== Array 1d: loop in array =====")
	for i := 0; i < 5; i++ {
		fmt.Println(i, arr[i])
	}

	// Declare a array 2d
	var arr2d [2][4]int

	fmt.Println("===== Array 2d: loop in array2d =====")
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d", arr2d[i][j])
		}
		fmt.Printf("\n")
	}

	// Declare a map
	var oldMap map[int]string
	oldMap = make(map[int]string)
	oldMap[1] = "Paris"
	oldMap[2] = "Rome"
	oldMap[3] = "Tokyo"
	oldMap[4] = "New Delhi"

	// Condition loop 1: Create a contion loop with omitting the increment statement
	fmt.Println("===== Condition loop 1: contion loop with omitting the increment statement < for ;i<=20; =====")
	// i := 2
	// for ;i <= 20; {
	// 	fmt.Printf("%d ", i)
	// 	i *= 2
	// }

	// Condition loop 2: Create a contion loop using for
	fmt.Println("===== Condition loop 2: make a simple condition loop: <for i <= 20> =====")
	i := 2
	for i <= 20 {
		fmt.Printf("%d\n", i)
		i *= 2
	}

	// Method 1: If you're looping over an array, slice, string, or map,
	// or reading from a channel, a range clause can manage the loop.
	fmt.Println("===== Method 1: loop with key and value =====")
	for key, name := range oldMap {
		fmt.Printf("Key %d, Name: %s\n", key, name)
	}

	// Method 2: If you only need the first item in the range (the key or index),
	// drop the second:
	fmt.Println("===== Method 2: loop with key =====")
	for key := range oldMap {
		fmt.Printf("Key %d, Name: %s\n", key, oldMap[key])
	}

	// Method 3: If you only need the second item in the range (the value),
	// use the blank identifier, an underscore, to discard the first:
	fmt.Println("===== Method 3: loop with value =====")
	for _, name := range oldMap {
		fmt.Printf("Name: %s\n", name)
	}

	// Method 4: We aslo loop if you know about len of slices
	fmt.Println("===== Method 4: loop with len =====")
	for i := 0; i < len(oldMap); i++ {
		fmt.Printf("i:  %d; Size of len: %d\n", i, len(oldMap))
		fmt.Printf("Name: %s\n", oldMap[i])
	}

	fmt.Println("=====Do you want to get a final exercise is infinite loop: [Y/N] ===== ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	x := input.Text()

	if x == "Y" {
		fmt.Println("Please enter: Ctrl + C to exit")
		for {

		}
	} else {
		fmt.Printf("You're done with all exercise")
	}
}
