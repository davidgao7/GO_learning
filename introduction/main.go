package main // special package name
import "fmt" // format

// create function
func main() { // main package need main function
	// code
	fmt.Println("Hello, World!")        // print
	fmt.Printf("Hello, %s!\n", "World") // print with format
	var intNum int64 = 10               // variable
	fmt.Println(intNum)                 // print 10
	var floatNum float64 = 10.5         // variable
	fmt.Println(floatNum)               // print 10.5
	// NOTE: string length are counted in ascii
	fmt.Println(len("test")) // print 4
	fmt.Println(len("A"))    // print 1
	fmt.Println(len("γ"))    // print 2

	// scan input
	var appleNum string
	fmt.Printf("number of apple?")
	fmt.Scanln(&appleNum) // takes address of appleNum
	fmt.Printf(appleNum)
}
