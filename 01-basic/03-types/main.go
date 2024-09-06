package main

import "fmt"

func main() {
	var my_str string = "Hello, Go"
	you_str := "Go is the best"
	var my_bool bool = true

	fmt.Printf("Type: %T, Value: %v\n", my_str, my_str)
	fmt.Printf("Type: %T, Value: %v\n", you_str, you_str)
	fmt.Printf("Type: %T, Value: %v\n", my_bool, my_bool)
}
