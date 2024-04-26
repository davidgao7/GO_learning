package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter an animal (cow, bird, snake) and an action (eat, move, speak) (e.g. 'cow eat'), q to quit:")
		fmt.Print("> ")
		line, _ := reader.ReadString('\n') // read a line from the user
		fields := strings.Fields(line)     // splits the line into words

		// fmt.Printf("fields: %v\n", fields)
		if len(fields) == 1 && fields[0] == "q" {
			// Exit the program
			fmt.Println("bye")
			os.Exit(0)
		}

		if len(fields) != 2 {
			fmt.Println("Invalid input. Please enter an animal and an action.")
			continue
		}

		animal, ok := animals[fields[0]]
		if !ok {
			fmt.Println("Invalid animal. Please enter 'cow', 'bird', or 'snake'.")
			continue
		}

		switch fields[1] {
		case "eat":
			animal.Eat()
			fmt.Println("")
			break
		case "move":
			animal.Move()
			fmt.Println("")
			break
		case "speak":
			animal.Speak()
			fmt.Println("")
			break
		default:
			fmt.Println("Invalid action. Please enter 'eat', 'move', 'speak' or 'q'.")
		}
	}
}
