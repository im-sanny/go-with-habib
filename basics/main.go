package main

import "fmt"

func legalAge() {
	age := 20

	if age >= 18 {
		fmt.Println("Your're eligible to marriage!")
	} else if age < 18 {
		fmt.Println("You're not eligible for marriage!")
	} else {
		fmt.Println("Grow up kiddo!")
	}

	switch {
	case age == 15 || age == 16:
		fmt.Println("Invalid")
	case age == 20 || age >= 18:
		fmt.Println("valid")
	default:
		fmt.Println("Bruh!")
	}
}

func main() {
	/* int, string, bool, float */
	a := 1
	a = 2

	b := true
	b = false

	c := "hello"
	c = "Hello world!"

	fmt.Println(a, b, c)
	legalAge()
}
