package main

import (
	"fmt"
)

func main() {

	var h int
	result := 1

	fmt.Println("enter a number")
	fmt.Scanln(&h)
	for i := 1; i <= h; i++ {

		result *= i
	}
	fmt.Printf("The factorial of %d is: %d\n", h, result)

}
