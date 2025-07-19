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

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func getTotal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

func main() {
	/* int, string, bool, float */
	a := 1
	a = 2

	b := true
	b = false

	c := "hello"
	c = "Hello world!"

	sum := add(1, 2)
	p, q := getTotal(20, 300)

	legalAge()
	fmt.Println(a, b, c)
	fmt.Println(sum)
	fmt.Println(p, q)
}
