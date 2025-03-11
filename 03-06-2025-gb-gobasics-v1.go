package main

import "fmt"

func main() {

	myslice1 := make([]int{})
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice2 := make([]int, 6)

	fmt.Printf("myslice2 = %v\n", myslice2)

	fmt.Printf("length = %d\n", len(myslice2))

	fmt.Printf("capacity = %d\n", cap(myslice2))

}
