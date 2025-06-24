package main

import "fmt"

func main() {
	x := 60
	//pointer of x
	var ptr *int = &x
	fmt.Println("Value of X: ", x)
	fmt.Println("Memory address of x: ", &x)
	fmt.Println("Pointer ptr: ", ptr)
	fmt.Println("Value at Pointer ptr: ", *ptr)
	//modify the value using the pointer
	*ptr = 100
	fmt.Println("New value of x", x)
}
