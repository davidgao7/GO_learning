package main

import (
	"fmt"
)

func main() {
	// find the character 'i', 'a', 'n'
	// eg. "iaaaaan" => "Found!"
	var input string
	fmt.Println("Enter a string:")

	fmt.Scan(&input)

	/*
	* enter a string that starts with the character ‘i’, ends with the character ‘n’,
	* and contains the character ‘a’, such as "iaaaan". Does the program print "Found!"*/
	if input[0] != 'i' || input[len(input)-1] != 'n' {
		fmt.Println("Not Found!")
		return
	}
	for i := 0; i < len(input); i++ {
		if input[i] == 'a' {
			fmt.Println("Found!")
			return
		}
	}
	fmt.Println("Not Found!")
}
