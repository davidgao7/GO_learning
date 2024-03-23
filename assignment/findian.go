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

	for i := 0; i < len(input); i++ {
		if input[i] == 'i' || input[i] == 'a' || input[i] == 'n' {
			fmt.Println("Found!")
			return
		}
	}
	fmt.Println("Not Found!")
}

func findc(input *string, c string) {
	for i := 0; i < len(*input); i++ {
		if string((*input)[i]) == c {
			fmt.Println("Found!")
			return
		}
	}
}
