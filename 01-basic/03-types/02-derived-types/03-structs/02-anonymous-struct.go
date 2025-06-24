package main

import "fmt"

func main(){
	// create anonymous struct & initialize
	mango := struct{
		Name string
		Size string
	}{
		Name: "Fazly",
		Size: "Big",
	}
	fmt.Println("Name: ", mango.Name)
	fmt.Println("Size: ", mango.Size)
}
