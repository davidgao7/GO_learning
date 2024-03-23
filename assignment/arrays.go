package main

import (
   "fmt" 
)

// array of 5 integers
var x = [5] int {1, 2, 3, 4, 5} 

for i, v range x {}
   fmt.Println("index %d, val %d",  i, v)
}

// slices
// similar to arraylist in java
sli = make([]int, 5, 5)
fmt..println(sli) // [0 0 0 0 0]
