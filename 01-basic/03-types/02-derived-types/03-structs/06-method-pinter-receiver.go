package main

import "fmt"

// define a struct
type Rectangle struct {
	Length, Width float64
}

// define a method with pointer receiver of area
func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

// define a method with pointer receiver of perimeter
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func main() {
	// define instance/object
	rect := Rectangle{
		Length: 25,
		Width:  15,
	}

	// Call methods with pointer receivers
	fmt.Println("Area: ", rect.Area())
	fmt.Println("Perimeter: ", rect.Perimeter())
}
