// create a function to perform bubble sort
package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

// take a slice of integers as an argument,
func BubbleSort(arr []int) []int {
	// and sort the slice in place using the bubble sort algorithm.
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

func main() {
	// take a input from user
	// The program should prompt the user to type in a sequence of up to 10 integers.
	fmt.Println("Enter a sequence of up to 10 integers:")
	var input int
	var arr []int

	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		arr = append(arr, input)
		if len(arr) == 10 {
			break
		}
	}
	// The program should write the integers out on one line, in sorted order,
	fmt.Println("Original integers:")
	fmt.Println(arr)
	fmt.Println("Sorted integers:")
	fmt.Println(BubbleSort(arr))
}
