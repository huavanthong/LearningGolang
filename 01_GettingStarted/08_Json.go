package main

import "fmt"

type Genre struct {
	Country string `json:"country"`
	Rock    string `json:"rock"`
}

type Music struct {
	Genre Genre `json:"genre"`
}

type ManyMusic struct {
	Genre []Genre `json:"genre"`
}

func main() {
	resp := Music{
		Genre{
			Country: "Taylor Swift",
			Rock:    "Aimee",
		},
	}

	fmt.Println("Details of the Genre")
	fmt.Printf("Check: %v \n", resp)

	resp2 := ManyMusic{
		[]Genre{
			{
				Country: "Taylor Swift",
				Rock:    "Aimee",
			},
			{
				Country: "Hua Van Thong",
				Rock:    "Giao Linh",
			},
		},
	}

	fmt.Println("Details of the ManyMusic")
	fmt.Printf("Check: %v \n", resp2)
}
