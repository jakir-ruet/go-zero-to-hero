package main

import "fmt"

// global variable
// var a int = 50
// var b int = 5
var a int
var b int

var addition int
var subtraction int
var multiplication int
var division int
var modulus int

func ArithOperation() {
	// local variable
	fmt.Print("Enter First Number ")
	fmt.Scan(&a)
	fmt.Print("Enter Second Number ")
	fmt.Scan(&b)
	addition = a + b
	subtraction = a - b
	multiplication = a * b
	division = a / b
	modulus = a % b
}
func main() {
	ArithOperation()
	fmt.Println(addition, subtraction, multiplication, division)
	fmt.Printf("Addition is %v \n Subtraction is %v \n Multiplication is %v \n Division is %v \n Reminder is %v \n", addition, subtraction, multiplication, division, modulus)
}