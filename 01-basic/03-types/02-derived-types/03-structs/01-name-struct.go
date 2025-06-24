package main

import "fmt"

// define struct named `myStruct`
type myStruct struct {
	Name    string
	PhoneNo string
	Address string
}

func main() {
	// define instance
	mystruct := myStruct{
		Name:    "Md Jakaria",
		PhoneNo: "01788916805",
		Address: "Rajshahi, Bangladesh",
	}
	fmt.Println("Name: ", mystruct.Name)
	fmt.Println("Phone No: ", mystruct.PhoneNo)
	fmt.Println("Address: ", mystruct.Address)
}
