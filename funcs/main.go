package main

import "fmt"

var a = 10

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	add(2, 3)
}

func init() {
	// fmt.Println("I'm a func that executes first!")
	fmt.Println(a)
	a = 20
}

// standard func: function with name are standard or named function
// init func: u can't call init func, program will call it by it's own, init function will be executed before any other function
//
