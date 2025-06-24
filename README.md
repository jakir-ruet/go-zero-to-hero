## More About Me – [Take a Look!](http://www.mjakaria.me)

## Welcome to Go Programming

Go is a high-level general purpose programming language that is statically typed and compiled. It is known for the simplicity of its syntax and the efficiency of development that it enables by the inclusion of a large standard library supplying many needs for common projects.

### Array vs. Slice

- **Array:** Fixed size; size is part of the type.
- **Slice:** Dynamic size; can grow or shrink. Slices are much more flexible than arrays.

```bash
func main() {
	var arr1 = [5]int{2, 3, 4, 5, 3}
	var inferredLength = [...]int{2, 3, 4, 5, 3}
	arr2 := [3]int{7, 8, 9}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(inferredLength)

	// accessing to specific element
	fmt.Println("Third Element ", arr1[2])
	// print index 1 to 4
	fmt.Println(arr1[1:4])

	// or using slice
	slice := arr1[1:4]
	fmt.Println(slice)
}
```

### Struct (Structure)

A struct (short for "structure") in Go is a composite data type that groups together zero or more fields of different types under a single name. Structs allow you to create custom data types that model real-world entities or logical data groupings, such as a Person, Car, or Employee.

In many ways, a struct in Go works like a class in object-oriented languages such as Python, Java, or C++, in that it holds both data (fields) and can be associated with methods. However, Go does not support traditional classes or inheritance. Instead, it promotes composition and interfaces for code reuse and polymorphism.

#### How many ways its use

1. Named Struct

```bash
type User struct {
    Name string
    Age  int
}
```

2. Anonymous Struct

```bash
person := struct {
    Name string
    Age  int
}{
    Name: "Eve",
    Age:  22,
}
```

3. Embedded Struct (Composition/Inheritance)

```bash
type Address struct {
    City string
    Zip  string
}

type Employee struct {
    Person
    Address
    Position string
}
```

- Access is flattened

```bash
e := Employee{
    Person:  Person{Name: "John", Age: 30},
    Address: Address{City: "NYC", Zip: "10001"},
    Position: "Manager",
}
```

4. Tagging Struct Fields (for JSON, DB, etc.)

```bash
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

5. Empty Struct

```bash
var s struct{} # Useful as a set placeholder
```

6. Method with Value Receiver

```bash
func (u User) Greet() {
    fmt.Println("Hi,", u.Name)
}
```

7. Method with Pointer Receiver

```bash
func (u *User) Birthday() {
    u.Age++
}
```

8. Method on Non-Struct Types (Custom Types)

```bash
type MyInt int

func (m MyInt) IsPositive() bool {
    return m > 0
}
```

9. Function that Takes Struct (not a method)

```bash
func Greet(u User) {
    fmt.Println("Hello", u.Name)
}
```

10. Interface Implementation

```bash
type Speaker interface {
    Speak()
}

func (u User) Speak() {
    fmt.Println("I'm", u.Name)
}
```

11. Factory Function for Struct

```bash
func NewUser(name string) User {
    return User{Name: name, Age: 0}
}
```

12. Array / Slice of Structs

```bash
people := []Person{
    {Name: "Tom", Age: 20},
    {Name: "Jerry", Age: 22},
}
```

13. Structs as Return Values or Parameters

```bash
func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}
```

- Or passed as parameters

```bash
func PrintPerson(p Person) {
    fmt.Println(p.Name, p.Age)
}
```

14. Struct in Maps or Channels

```bash
type Event struct {
    ID   int
    Name string
}

eventMap := map[int]Event{}
eventChan := make(chan Event)
```

15. Method Chaining with Structs

```bash
type Builder struct {
    name string
}

func (b *Builder) SetName(name string) *Builder {
    b.name = name
    return b
}
```

16. Function Fields in Structs

```bash
type Processor struct {
    Handle func(data string) string
}
```

### Interface

An interface is a collection of method signatures that a type must implement to be considered an implementation of that interface. Also, it's a type that defines a set of method signatures. A type satisfies an interface by implementing all the methods declared in that interface. It's enable `polymorphism` and `abstraction`.

### Concurrency

It's the ability of a program to make progress on multiple tasks at the same time. Go has first-class support for concurrency using:

- Goroutines
- Channels
- select statements
- sync package (mutexes, waitgroups, etc.)

#### Goroutine

A goroutine is a lightweight thread managed by the Go runtime. It allows you to run functions concurrently using the go keyword.

#### Channel

A channel is a typed pipe that allows goroutines to communicate safely and synchronize without shared memory.

#### select Statement

Select lets you wait on multiple channel operations. It’s like a switch for channels.

## With Regards, `Jakir`

[![LinkedIn][linkedin-shield-jakir]][linkedin-url-jakir]
[![Facebook-Page][facebook-shield-jakir]][facebook-url-jakir]
[![Youtube][youtube-shield-jakir]][youtube-url-jakir]

### Wishing you a wonderful day! Keep in touch

<!-- Personal profile -->

[linkedin-shield-jakir]: https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white
[linkedin-url-jakir]: https://www.linkedin.com/in/jakir-ruet/
[facebook-shield-jakir]: https://img.shields.io/badge/Facebook-%231877F2.svg?style=for-the-badge&logo=Facebook&logoColor=white
[facebook-url-jakir]: https://www.facebook.com/jakir.ruet/
[youtube-shield-jakir]: https://img.shields.io/badge/YouTube-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white
[youtube-url-jakir]: https://www.youtube.com/@mjakaria-ruet/featured
