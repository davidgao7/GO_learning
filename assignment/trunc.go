package main

import (
	"fmt"
	"math"
)

func main() {
	var floatNum float64 = 123.456789
	fmt.Println("Original floating point number:", floatNum)

	// Truncate the floating point number
	truncatedNum := math.Trunc(floatNum)
	fmt.Println("Truncated floating point number:", truncatedNum)
}
