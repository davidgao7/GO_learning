/*
Write a program which allows the user to create a set of animals and to get information
about those animals. Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user. Note that the user can
define animals of a chosen type, but the types of animals are restricted to either cow,
bird, or snake. The following table contains the three types of animals and their associated data.

Your program should present the user with a prompt, “>”, to indicate that the user can type a
request. Your program should accept one command at a time from the user, print out a response,
and print out a new prompt on a new line. Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings.
The first string is “newanimal”. The second string is an arbitrary string which will be the name
of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!”
on the screen.

Each “query” command must be a single line containing 3 strings.
The first string is “query”. The second string is the name of the animal.
The third string is the name of the information requested about the animal,
either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird,
and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the
appropriate type. Your program should call the appropriate method when the user issues a query command.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	Food, Locomotion, Noise, Name string
}

func (a Animal) Eat() {
	fmt.Println(a.Food)
}

func (a Animal) Move() {
	fmt.Println(a.Locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.Noise)
}

func main() {

	cows := []Animal{}
	birds := []Animal{}
	snakes := []Animal{}

	// initialize the map with the 3 types of animals
	animals := map[string][]Animal{
		"cow":   cows,
		"bird":  birds,
		"snake": snakes,
	}

	reader := bufio.NewReader(os.Stdin)
	message := `Enter "newanimal" + "animal name" +"animal type" for creating animal 
    or 
    "query" + "animal name" + "eat" or "move" or "speak", separate token with space, q to quit:
`

	for {
		fmt.Println(message)
		fmt.Print("> ")
		line, _ := reader.ReadString('\n') // read a line from the user
		fields := strings.Fields(line)     // splits the line into words, all input is 3 words, so lets split at the very beginning

		// fmt.Printf("fields: %v\n", fields)
		if len(fields) == 1 && fields[0] == "q" {
			// Exit the program
			fmt.Println("bye")
			os.Exit(0)
		}

		if len(fields) != 3 {
			fmt.Print(fields)
			fmt.Println("Invalid input.")
			continue
		}

		command := fields[0]
		switch command {

		case "newanimal":
			/*
			 * Each “newanimal” command must be a single line containing three strings.
			 * The first string is “newanimal”. The second string is an arbitrary string which will be
			 * the name of the new animal. The third string is the type of the new animal,
			 * either “cow”, “bird”, or “snake”.  Your program should process each newanimal
			 * command by creating the new animal and printing “Created it!” on the screen.
			 * */
			animalType := fields[2]
			animalName := fields[1]

			switch animalType {
			case "cow":
				animals[animalType] = append(animals[animalType], Animal{Food: "grass", Locomotion: "walk", Noise: "moo", Name: animalName})
			case "bird":
				animals[animalType] = append(animals[animalType], Animal{Food: "worms", Locomotion: "fly", Noise: "peep", Name: animalName})
			case "snake":
				animals[animalType] = append(animals[animalType], Animal{Food: "mice", Locomotion: "slither", Noise: "hsss", Name: animalName})
			default:
				fmt.Println("Invalid animal type. Please use 'cow', 'bird', or 'snake'.")
			}

			fmt.Println("Created it!")
			break
		case "query":
			/*
			 * Each “query” command must be a single line containing 3 strings. The first string is
			 * “query”. The second string is the name of the animal.
			 * The third string is the name of the information requested about the animal, either
			 * “eat”, “move”, or “speak”. Your program should process each query
			 * command by printing out the requested data.
			 **/
			animalName := fields[1] // name of the animal
			action := fields[2]     // eat | move | speak

			switch action {
			case "eat":
				for animalType, animal_list := range animals {
					for _, animal := range animal_list {
						if animal.Name == animalName {
							fmt.Printf("%s %s eat ", animalType, animalName)
							animal.Eat()
						}
					}
				}
				break
			case "move":
				for animalType, animal_list := range animals {
					for _, animal := range animal_list {
						if animal.Name == animalName {
							fmt.Printf("%s %s ", animalType, animalName)
							animal.Move()
						}
					}
				}
				break
			case "speak":
				for animalType, animal_list := range animals {
					for _, animal := range animal_list {
						if animal.Name == animalName {
							fmt.Printf("%s %s ", animalType, animalName)
							animal.Speak()
						}
					}
				}
				break
			default:
				fmt.Println("Invalid action type. Please use 'eat', 'move', or 'speak'.")
			}
			break
		default:
			fmt.Println("Invalid command. Please enter 'newanimal' or 'query'.")
		}
	}
}
