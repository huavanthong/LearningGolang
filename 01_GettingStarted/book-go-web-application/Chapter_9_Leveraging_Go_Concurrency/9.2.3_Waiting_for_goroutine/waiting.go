package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	// Step 3: Decrement the counter using the Done method.
	wg.Done()
}
func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	// Step 3: Decrement the counter using the Done method.
	wg.Done()
}
func main() {
	//Step 1: Declare a variable from sync.WaitGroup 
	var wg sync.WaitGroup
	//Step 2: Set up the WaitGroup's counter using the Add method.
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	// Step 4: Call the Wait method, which will block until the counter is 0
	wg.Wait()
}
