package main

import "fmt"

// passing array to function like parameter
func arrayPassFunction(arr *[5]int) {
	arr[0] = 99
}

func main() {
	nums := [5]int{3, 2, 5, 6, 9}
	fmt.Println("Before function call:", nums)
	arrayPassFunction(&nums)
	// replace 3 by 99
	fmt.Println("After function call:", nums)
}
