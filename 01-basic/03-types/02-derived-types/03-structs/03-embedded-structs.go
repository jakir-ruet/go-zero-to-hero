package main

import "fmt"

type Person struct{
	Name string
	School string
}
type Student struct{
	Person
	Class string
}

func main(){
	// defining instance
	student := Student{
		Person: Person{
			Name: "Md Jakaria",
			School: "ABC High School",
		},
		Class: "Nine",
	}
	fmt.Println("Student name: ", student.Name)
	fmt.Println("School: ", student.School)
	fmt.Println("Class: ", student.Class)
}
