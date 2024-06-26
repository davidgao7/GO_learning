// Let us assume the following formula for
// displacement s as a function of time t, acceleration a, initial velocity vo,
// and initial displacement so.
//
// s = ½ a t2 + vot + so
//
// Write a program which first prompts the user
// to enter values for acceleration, initial velocity, and initial displacement.
// Then the program should prompt the user to enter a value for time and the
// program should compute the displacement after the entered time.
//
// You will need to define and use a function
// called GenDisplaceFn() which takes three float64
// arguments, acceleration a, initial velocity vo, and initial
// displacement so. GenDisplaceFn()
// should return a function which computes displacement as a function of time,
// assuming the given values acceleration, initial velocity, and initial
// displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
// float64 argument which is the displacement travelled after time t.
//
// For example, let’s say that I want to assume
// the following values for acceleration, initial velocity, and initial
// displacement: a = 10, vo = 2, so = 1. I can use the
// following statement to call GenDisplaceFn() to
// generate a function fn which will compute displacement as a function of time.
//
// fn := GenDisplaceFn(10, 2, 1)
//
// Then I can use the following statement to
// print the displacement after 3 seconds.
//
// fmt.Println(fn(3))
//
// And I can use the following statement to print
// the displacement after 5 seconds.
//
// fmt.Println(fn(5))

// Submit your Go program source code.

package main

import (
	"fmt"
)

// test two different sets of values for acceleration, initial velocity, and initial displacement
// and time
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	// test 1
	// fn := GenDisplaceFn(2, 10, 1)
	// fmt.Println(fn(3)) // 46
	// fmt.Println(fn(5)) // 156
	// // test 2
	// fn = GenDisplaceFn(5, 3, 2)
	// fmt.Println(fn(3)) // 38.5
	// fmt.Println(fn(5)) // 72.5
	/*
		* NOTE:
			* The type of data it tries to read is determined by the type of the variable.
			* If the user's input can't be converted to the variable's type,
		* `fmt.Scan()` will return an error.
	*/
	var a, vo, so, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scanln(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scanln(&vo)
	fmt.Print("Enter initial displacement: ")
	fmt.Scanln(&so)
	fmt.Print("Enter time: ")
	fmt.Scanln(&t)

	result := GenDisplaceFn(a, vo, so)(t)
	fmt.Printf("Displacement after %.2f seconds: %.2f with velocity %.2f\n", t, result, vo+a*t)

}
