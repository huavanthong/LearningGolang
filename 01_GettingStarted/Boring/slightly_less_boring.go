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
	slightly_boring("boring !")
}
