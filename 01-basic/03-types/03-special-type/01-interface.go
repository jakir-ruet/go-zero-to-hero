package main

import "fmt"

// define an Greeter interface
type Greeter interface {
	Greet() string
}

// define English struct for English speaker
type English struct {
	Name string
}

func (English) Greet() string {
	return "Hello, I speaking in English"
}

// define Spanish struct for Spanish speaker
type Spanish struct {
	Name string
}

func (Spanish) Greet() string {
	return "Hello, I speaking in Spanish"
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	sayHello(English{})
	sayHello(Spanish{})
}
