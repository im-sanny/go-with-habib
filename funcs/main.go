package main

import "fmt"

var a = 10 //expression

// standard or named function
// func add(a, b int) {
// 	fmt.Println(a + b)
// }

func main() {
	// add(2, 3)

	// anonymous func & IIFE
	// func(a int, b int) {
	// 	sum := a + b
	// 	fmt.Println(sum)
	// }(1, 2)

	// func expression
	add := func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}
	add(4, 5)
}

func init() {
	fmt.Println("I'll execute first!")
	fmt.Println(a)
	a = 20
}

// standard func: function with name are standard or named function
// init func: u can't call init func, program will call it by it's own, init function will be executed before any other function
// anonymous func: functions without name is anonymous func.
// IIFE: if a function call/invoked immediately after writing then it's IIFE.
// if we assign a function in a variable then it'll be function expression.
// In local scope we can't invoke or call a function literal (function expression) before its assignment. We must invoke it from a lower position, after the function expression is defined.
