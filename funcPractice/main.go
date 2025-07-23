package main

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome!")
}

func getUsername() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getNumber() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("enter second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func display(name string, sum int) {
	fmt.Println("hello ", name)
	fmt.Println("total ", sum)
}

func bye() {
	fmt.Println("good bye")
}

func main() {
	welcomeMessage()
	name := getUsername()
	num1, num2 := getNumber()
	sum := add(num1, num2)
	display(name, sum)
	bye()
}
