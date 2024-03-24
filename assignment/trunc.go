package main

import (
	"fmt"
	"math"
)

func main() {
	var floatNum float64
	fmt.Println("Enter a floating point number:")
	// fmt.Scan(&floatNum)
	fmt.Scanf("%f", &floatNum)

	fmt.Println("Original floating point number:", floatNum)

	// Truncate the floating point number
	truncatedNum := math.Trunc(floatNum)
	fmt.Println("Truncated floating point number:", truncatedNum)
}
