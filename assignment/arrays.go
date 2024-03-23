package main

import (
	"fmt"
)

func main() {
	// array of 5 integers
	var x = [5]int{1, 2, 3, 4, 5}

	for i, v := range x {
		fmt.Printf("index %d, val %d\n", i, v)
	}

	// slices
	// similar to arraylist in java
	sli := make([]int, 5, 5)
	fmt.Println(sli) // [0 0 0 0 0]
	// the slice can automatically grow when you append
	slice := append(sli, 6) // [0 0 0 0 0 6]
	fmt.Println(slice)      // [0 0 0 0 0 6]

}
