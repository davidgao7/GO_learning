// The goal of this assignment is to practice working with the ioutil and os packages in Go.

package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Print("Enter the name of the text file: ")
	fmt.Scan(&fileName)

	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var names []Name
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}
		names = append(names, Name{fname: parts[0], lname: parts[1]})
	}

	for _, name := range names {
		fmt.Println("First Name:", name.fname, ", Last Name:", name.lname)
	}
}
