package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("an error occurred")
}

func main() {
	err := doSomething()
	if err != nil {
		err = errors.WithStack(err)
		fmt.Println(err)
	}
}
