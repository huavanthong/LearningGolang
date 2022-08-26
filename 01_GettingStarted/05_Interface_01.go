package main

import (
	"fmt"
)

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	vegetableSet := map[string]bool{
		"potato":  true,
		"cabbage": true,
		"carrot":  true,
	}

	fruitRank := map[int]string{
		1: "strawberry",
		2: "raspberry",
		3: "blueberry",
	}

	fmt.Printf("vegetableSet keys: %+v\n", keys(vegetableSet))
	fmt.Printf("fruitRank keys: %+v\n", keys(fruitRank))
}
