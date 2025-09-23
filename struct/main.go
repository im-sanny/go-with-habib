package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func person() {
	var user1 User

	user1 = User{
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

// struct: in golang struct is method for creating custom data types
