package main

import "fmt"

func main() {
	// Bool Type
	fruit := true
	trees := false
	fmt.Println(fruit)
	fmt.Println(trees)
	// int, int8, int16, int32, int64
	var x int = 42      // Default int (32 or 64 bits)
	var y int8 = -128   // 8-bit signed integer
	var z int16 = 32000 // 16-bit signed integer

	fmt.Println(x) // Output: 42
	fmt.Println(y) // Output: -128
	fmt.Println(z) // Output: 32000

	// uint, uint8, uint16, uint32, uint64
	var a uint = 42      // Default unsigned integer (32 or 64 bits)
	var b uint8 = 255    // 8-bit unsigned integer
	var c uint16 = 65535 // 16-bit unsigned integer

	fmt.Println(a) // Output: 42
	fmt.Println(b) // Output: 255
	fmt.Println(c) // Output: 65535

	// float32, float64
	var pi float64 = 3.14159
	var e float32 = 2.71828

	fmt.Println(pi) // Output: 3.14159
	fmt.Println(e)  // Output: 2.71828

	// string
	var greeting string = "Hello, Go!"
	fmt.Println(greeting) // Output: Hello, Go!

	// Rune (rune)
	var char rune = 'A' // Single character
	fmt.Println(char)   // Output: 65 (Unicode value of 'A')

	// Byte (byte)
	var letter byte = 'A' // 'A' in byte format
	fmt.Println(letter)   // Output: 65 (Unicode value of 'A')

	// Creating complex numbers
    var c1 complex64 = complex(3.0, 4.0)   // complex64 with float32 parts
    var c2 complex128 = complex(1.0, 2.0)  // complex128 with float64 parts

    fmt.Println("Complex number 1:", c1)   // Output: (3+4i)
    fmt.Println("Complex number 2:", c2)   // Output: (1+2i)

    // Extracting the real and imaginary parts of a complex number
    realPart := real(c1)
    imaginaryPart := imaginary(c1)

    fmt.Println("Real part of c1:", realPart)          // Output: 3
    fmt.Println("Imaginary part of c1:", imaginaryPart) // Output: 4
}
