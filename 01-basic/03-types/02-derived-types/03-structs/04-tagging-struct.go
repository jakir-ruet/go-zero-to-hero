package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"` // Capitalized to export the field
}

func main() {
	// defining instance
	p := Person{
		Name:    "Md Jakaria",
		Age:     30,
		Email:   "",
		IsAdmin: true, // Use IsAdmin instead of isAdmin
	}

	// Marshalling to JSON
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error Handling: ", err)
		return
	}

	// Printing the JSON output
	fmt.Println(string(data)) // {"name":"Md Jakaria","age":30,"email":"","is_admin":true}
}
