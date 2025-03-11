package main

import "fmt"

func main() {

	myslice := []int{1, 2, 8, 3, 6, 6}
	/*myslice := myarray[2:4]

	fmt.Println(myslice)
	fmt.Printf("Length of myslice = %d\n", len(myslice))
	fmt.Printf("Capacity of myslice=%d\n", cap(myslice))*/
	myslice[0] = 5

	fmt.Println(myslice)

}
