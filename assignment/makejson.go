package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	// make a dictionary holder
	data := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	fullName, _ := reader.ReadString('\n')
	// get rid of the newline character
	fullName = strings.TrimSuffix(fullName, "\n")

	fmt.Print("Enter your address: ")
	address, _ := reader.ReadString('\n')
	// get rid of the newline character
	address = strings.TrimSuffix(address, "\n")

	data["name"] = fullName
	data["address"] = address

	bytes, _ := json.Marshal(data)
	fmt.Printf("JSON obj: %v", string(bytes)) // %v for any type
}
