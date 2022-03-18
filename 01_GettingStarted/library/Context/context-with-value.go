package main

import (
	"context"
	"fmt"
)

func main() {
	// Step 1: Create a context in backgroud
	ctx := context.Background()
	// Step 2: Call function to add a value
	ctx = addValue(ctx)
	// Step 3: Call function to read that value.
	readValue(ctx)
}

func addValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "key", "test-value")
}

func readValue(ctx context.Context) {
	val := ctx.Value("key")
	fmt.Println(val)
}
