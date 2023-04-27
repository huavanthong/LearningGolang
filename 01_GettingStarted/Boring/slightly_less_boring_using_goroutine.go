package main

import (
	"fmt"
	"math/rand"
	"time"
)

func slightly_boring(msg string) {

	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// The go statement runs the function as usual, but doesn't make the caller wait.
	go slightly_boring("boring !")

	fmt.Println("I'm listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring: I'm leaving")
}
