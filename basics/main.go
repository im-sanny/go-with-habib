package main

import "fmt"

// Check if age is eligible for marriage
func legalAge() {
	age := 20

	if age >= 18 {
		fmt.Println("You're eligible for marriage!")
	} else {
		fmt.Println("You're not eligible for marriage!")
	}

	switch {
	case age == 15 || age == 16:
		fmt.Println("Invalid")
	case age >= 18:
		fmt.Println("Valid")
	default:
		fmt.Println("Bruh!")
	}
}

// Adds two numbers and returns the result
func add(num1, num2 int) int {
	return num1 + num2
}

// Returns both sum and product of two numbers
func getTotal(num1, num2 int) (sum, mul int) {
	sum = num1 + num2
	mul = num1 * num2
	return
}

// Prints a welcome message
func sayHello(name string) {
	fmt.Println("Welcome to the course,", name)
}

func main() {
	// Variable declarations
	a := 2
	b := false
	c := "Hello world!"

	// Function usage
	sum := add(1, 2)
	total, product := getTotal(20, 300)

	sayHello("Sun")
	legalAge()

	// Output results
	fmt.Println(a, b, c)
	fmt.Println(sum)
	fmt.Println(total, product)
}
