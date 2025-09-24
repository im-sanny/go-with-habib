package main

import "fmt"

type User struct { // struct definition
	Name string			 // member variable or property
	Age  int
}

func person() {
	user1 := User{ 	// instance of User type
		Name: "Sun",
		Age:  69,
	}

	fmt.Println("Name:", user1.Name)
	fmt.Println("Age:", user1.Age)

	user2 := User{
		Name: "Bun",
		Age:  96,
	}

	fmt.Println("Name:", user2.Name)
	fmt.Println("Age:", user2.Age)

}

func main() {
	person()
}

/*
Struct: In Go, structs are used to define new types that contain multiple fields.

Instance: When I create a value of a custom type in Go, it's called a value or instance. Some also called it object but that's least preferable.

Instantiate: process of making instance called instantiate.

*/
