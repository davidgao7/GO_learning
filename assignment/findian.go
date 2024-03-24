package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// find the character 'i', 'a', 'n'
	// eg. "iaaaaan" => "Found!"
	var input string
	fmt.Println("Enter a string:")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	input1 := strings.ToLower(input)
	// fmt.Println(input1)
	/*
	* enter a string that starts with the character ‘i’, ends with the character ‘n’,
	* and contains the character ‘a’, such as "iaaaan". Does the program print "Found!"*/
	// fmt.Println(string(input1[0])) // this is byte, neet to convrt to string
	// fmt.Println(string(input1[len(input)-1]))
	if string(input1[0]) != "i" || string(input1[len(input)-1]) != "n" {
		fmt.Println("Not Found!")
		return
	}
	for i := 0; i < len(input1); i++ {
		// fmt.Println(input1[i])
		// fmt.Println("==================")
		if string(input1[i]) == "a" {
			fmt.Println("Found!")
			return
		}
	}
	fmt.Println("Not Found!")
}
