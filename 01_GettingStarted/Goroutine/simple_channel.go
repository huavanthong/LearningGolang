package main

import (
	"fmt"
	"time"
)

func main() {
	// start to measure the execution time of a program:
	start := time.Now()

	// Create a new channel with make
	c := make(chan int)

	// Launch a new goroutine
	go func() {
		// Send a value to the channel
		c <- 42
	}()

	// Read the value from the channel
	fmt.Println(<-c)

	// calculate the elapsed time and print it.
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}
