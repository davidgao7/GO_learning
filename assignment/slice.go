package main

import (
	"fmt"
	"strconv"
)

/*
* Write a program which prompts the user to enter integers and stores the integers in a sorted
* slice. The program should be written as a loop. Before entering the loop, the program
* should create an empty integer slice of size (length) 3. During each pass through the loop,
* the program prompts the user to enter an integer to be added to the slice. The program adds
* the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
* The slice must grow in size to accommodate any number of integers which the user decides to
* enter. The program should only quit (exiting the loop) when the user enters the character ‘X’
* instead of an integer.*/

func main() {
	slice := make([]int, 0, 3)
	var numStr string
	var num int
	for {
		fmt.Println("Enter an integer to be added to the slice:")
		_, err := fmt.Scanf("%s", &numStr)
		if numStr == "X" {
			return
		}
		num, err = strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}
		slice = append(slice, num)
		sort(slice)
		fmt.Println("Sorted slice:", slice)
	}
}

func sort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
