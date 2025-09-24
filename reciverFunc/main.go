package main

import "fmt"

type User struct { // struct definition
	Name string // member variable or property
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// receiver function
func (usr User) printDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

func main() {
	var user1 User
	user1 = User{
		Name: "Tom",
		Age:  69,
	}

	printUserDetails(user1)
}

/*
Receiver func: receiver function can be used when u have custom type or struct
*/
