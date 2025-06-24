package main

import (
	"fmt"
)

func main() {
	// Declaring an array of length 5
	var arr1 = [5]int{2, 3, 4, 5, 3}
	// Array with inferred length
	var inferredLength = [...]int{2, 3, 4, 5, 3}
	// Array of length 3
	arr2 := [3]int{7, 8, 9}

	// Print the arrays
	fmt.Println(arr1)           // [2 3 4 5 3]
	fmt.Println(arr2)           // [7 8 9]
	fmt.Println(inferredLength) // [2 3 4 5 3]

	// Accessing a specific element of the array
	fmt.Println("Third Element", arr1[2]) // Third element of arr1: 4

	// Printing a slice from index 1 to 4 (not including index 4)
	fmt.Println(arr1[1:4]) // [3 4 5]

	// Using slice as an explicit variable
	slice := arr1[1:4]
	fmt.Println(slice) // [3 4 5]
}
