package main

import "fmt"

// User struct definition
type User struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	// ------------------------------
	// Example 1: Struct pointer
	// ------------------------------
	sun := User{
		Name:   "Sun",
		Age:    69,
		Salary: 0,
	}

	p := &sun // p is a pointer to struct User

	// Access struct fields directly via pointer (Go does it automatically)
	fmt.Println(p.Age)

	// Dereferencing pointer: *p gives the actual struct value
	fmt.Println(*p)

	// ------------------------------
	// Example 2: Basic pointer with int
	// ------------------------------
	x := 20
	fmt.Println("x =", x)

	p2 := &x // &x = address of variable x
	*p2 = 30 // *p2 = value at that address → updates x

	fmt.Println("x =", x)
	fmt.Println("Address of x:", p2)
	fmt.Println("Value at the address:", *p2)

	// ------------------------------
	// Example 3: Pointer with array
	// ------------------------------
	arr := [3]int{1, 2, 3}
	printArray(&arr)
}

// Function to accept an array pointer
func printArray(numbers *[3]int) {
	fmt.Println(*numbers) // dereference pointer to get array values
}

/*
------------------------------
 Notes on Pointers in Go
------------------------------

1. Pointer:
   - A pointer stores the memory address of a variable.
   - Example: for a person, his home address is like his pointer.

2. &: (Address-of operator)
   - Used to get the address of a variable.
   - Example: p := &x → p stores the address of x.

3. *: (Dereference operator)
   - Used to access the value stored at a memory address.
   - Example: *p → value at address p.

4. Struct pointer:
   - In Go, if you have a pointer to a struct, you can access its fields directly (Go automatically handles dereferencing).
   - Example: p.Age works, no need to write (*p).Age.

5. Pass by value vs Pass by reference:
   - Pass by value: function receives a copy of the variable.
   - Pass by reference: function receives the address (pointer), so it can modify the original variable.

6. Important:
   - You can’t use * directly unless you have a pointer (created with &).
   - Pointers in Go point to memory in RAM, not to "hard disk".

*/
