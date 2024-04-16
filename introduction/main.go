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
	fmt.Print("\n=================")
	fmt.Println()

	// maps
	var m map[string]int
	m = make(map[string]int)
	m["key"] = 10
	fmt.Println(m["key"]) // print 10
	delete(m, "key")
	fmt.Println(m["key"]) // print 0

	// two value assignment for maps
	val, ok := m["key"]  // val is value, ok is presence of key
	fmt.Println(val, ok) // print 0, false

	// iterate through maps
	print("==========Iterate through maps==========\n")
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	for key, val := range m {
		fmt.Println(key, val)
	}
	print("==========Iterate through maps==========\n")

	x := []int{4, 8, 5}
	y := -1
	for _, elt := range x {
		if elt > y {
			y = elt
		}
	}
	fmt.Print(y)

	print("\n===================================\n")
	x = []int{4, 8, 5}
	Y := x[0:2]
	z := x[1:3]
	Y[0] = 1
	z[1] = 3
	fmt.Print(x)

	print("\n===================================\n")
	a := [...]int{1, 2, 3, 4, 5}
	b := a[0:2]
	c := a[1:4]
	fmt.Print(len(b), cap(b), len(c), cap(c))

	print("\n===================================\n")
	d := map[string]int{
		"ian": 1, "harris": 2}
	for i, j := range d {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}
	print("\n===================================\n")
	type P struct {
		x string
		y int
	}
	e := P{"x", -1}
	f := [...]P{P{"a", 10},
		P{"b", 2},
		P{"c", 3}}
	for _, z := range f {
		if z.y > e.y {
			e = z
		}
	}
	fmt.Println(e.x)
	print("\n===================================\n")
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))

	print("\n===================================\n")
	print(add(1, 2))
	print("\n===================================\n")

	// like python, you can assign more than 1 variable
	var1, var2 := returnMultiple(1) // print 1, return 1, 2
	println(var1, var2)             // print 1, 2
	print("===================================\n")

	// call by reference
	aa := 5
	reference(&aa) // change the value of aa at memory address from 5 to 10
	fmt.Print(aa)  // 10
	print("\n===================================\n")

}

// create a function
// func var can be simplify if have the same data type
func add(a, b int) int { // get two int type as inputs, return an int type
	fmt.Println(a + b)
	return a + b
}

func returnMultiple(i int) (int, int) {
	fmt.Println(i)
	return 1, 2
}

func reference(a *int) {
	*a = 10
}

// go : use slice instead of array
// create a slice
var cars = []string{"BMW", "Benz", "Audi"}
