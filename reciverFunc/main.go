package main

import "fmt"

type User struct { // struct definition
	Name string // property
	Age  int
}

// value receiver method (doesn't change original User)
func (usr User) printDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// another method with a parameter
func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	user1 := User{
		Name: "Tom",
		Age:  69,
	}

	user1.printDetails()
	user1.call(1)
}

/*
	- Receiver func: Receiver functions can be used when you have a custom type or struct. They only work with custom types, and use of receiver functions is suggested with structs.
	- Value receiver: gets a copy of the struct → changes inside don’t affect the original.
	- Pointer receiver: gets the actual struct’s address → can modify the original data.
	- Use receiver funcs when you want to give "behavior" to a struct.
*/
