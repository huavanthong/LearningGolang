/*

Number 1: 	9  		==> 1001 			==> gap of length = 2
Number 2: 	529		==> 1000010001		==> gap of length = 4
Number 3: 	20		==> 10100 			==> gap of length = 1
Number 4:	15		==> 1111 			==> gap of length = 0
Number 5: 	1041	==> 10000010001		==> gap of length = 5
Number 6: 	17		==> 10001 			==> gap of length = 3

*/
package main

import (
	"fmt"
)

func Solution(N int) int {
	var reverse_binary []int

	for N > 0 {
		reverse_binary = append(reverse_binary, N%2)
		N = N / 2
	}

	binary := []int{}
	for i := len(reverse_binary) - 1; i >= 0; i-- {
		binary = append(binary, reverse_binary[i])
	}

	for i := 0; i < len(binary); i++ {
		fmt.Printf("%d", binary[i])
	}
	fmt.Println("\nEnd method 1")

	for key := len(binary) - 1; key >= 0; key-- {
		fmt.Printf("%d", binary[key])
	}
	fmt.Println("\nEnd method 2")

	for key, _ := range binary {
		fmt.Printf("%d", binary[key])
	}
	fmt.Println("\nEnd method 3")
	fmt.Println("Calc binary gap: ... ")

	max := 0
	for i := 0; i < len(binary); i++ {
		count := 0
		if binary[i] == 1 {
			for j := i + 1; j < len(binary); j++ {
				if binary[j] == 0 {
					count++
				}

				if binary[j] == 1 {
					break
				}
			}
		}

		if count > max {
			max = count
		}
	}

	return max
}

func main() {

	fmt.Println("Value: ", Solution(529))

}
