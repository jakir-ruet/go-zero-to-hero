package main

import "fmt"

// // If, If else & If else if Section
// func IfStatement(){
// 	x := 15
// 	y := 3
// 	if x > y {
// 		// Correct use of fmt.Printf for formatted output
// 		fmt.Printf("%v is greater than %v, where x > y\n", x, y)
// 	}
// }

// func IfElseStatement(){
// 	time := 8
// 	if (time <= 8) {
// 		fmt.Println("Good Morning")
// 	}	else {
// 		fmt.Println("Good Evening")
// 	}
// }

// func IfElseIfStatement(){
// 	time := 23
// 	if (time <= 12) {
// 		fmt.Println("Good Morning")
// 	} else if (time >= 18 && time <= 22) {
// 		fmt.Println("Good Evening")
// 	}else {
// 		fmt.Println("Good Night")
// 	}
// }

// // For loop, while loop, do while loop
func ForLoop() {
    i := 10
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    for i := range 3 {
        fmt.Println("range", i)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
func main() {
	// // If, If else & If else if Section
	// IfStatement()
	// IfElseStatement()
	// IfElseIfStatement()
	// // For loop, while loop, do while loop
	ForLoop()
}

