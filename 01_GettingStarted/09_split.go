package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var x string
	var y string
	var w string
	var z string

	// Intializing the Strings
	x = "Hello! World"
	y = "This, is, a, sample, program"
	w = "1-2-3-4"
	z = "*model.product_phone)"

	// Display the Strings
	fmt.Println("String 1:", x)
	fmt.Println("String 2:", y)
	fmt.Println("String 3:", w)

	// Using the Split Function
	res1 := strings.Split(x, "!")
	res2 := strings.Split(y, ",")
	res3 := strings.Split(w, "-")
	res4 := strings.Split(y, "a")
	res5 := strings.Split(z, ".")

	// Display the Split Output
	fmt.Println("Split String 1:", res1)
	fmt.Println("Split String 2:", res2)
	fmt.Println("Split String 3:", res3)
	for key, value := range res3 {
		fmt.Println(key, value)
	}
	fmt.Println("Split String 2 with delimiter:", res4)
	fmt.Println("Split dot:", res5)
	fmt.Println("Split dot second: ", res5[1])
	fmt.Println("Split dot second: ", returnValue(z))

	fmt.Println("Get type of variable: ", TypeOf(z))

}

// Get type of variable
func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func returnValue(v interface{}) string {
	// example value: *models.product_phone
	// model := fmt.Sprintf("%T", v)
	model := "*model.product_phone"

	match, err := regexp.MatchString("[a-z].[a-z]", model)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println("Check model: ", model)
	typeModel := strings.Split(model, ".")

	fmt.Println("Check type model: ", typeModel)

	if len(typeModel) > 2 {
		return "Undefine"
	}

	return typeModel[1]
}
