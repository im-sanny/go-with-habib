package main

import "fmt"

// var a = 10 //expression

func processOperation(a int, b int, op func(x int, y int)) {
	op(a, b)
}

func addr(x int, y int) {
	z := x + y
	fmt.Println(z)
}

// // standard or named function
// func add(a int, b int) { //parameter => a, b
// 	fmt.Println(a + b)
// }

func main() {
	// add(2, 3) //argument => 2, 3

	processOperation(5, 6, addr)

	// anonymous func & IIFE
	// func(a int, b int) {
	// 	sum := a + b
	// 	fmt.Println(sum)
	// }(1, 2)

	// func expression
	// add := func(a int, b int) {
	// 	c := a + b
	// 	fmt.Println(c)
	// }
	// add(4, 5)
}

// func init() {
// 	fmt.Println("I'll execute first!")
// 	fmt.Println(a)
// 	a = 20
// }

// standard func: function with name are standard or named function
// init func: u can't call init func, program will call it by it's own, init function will be executed before any other function
// anonymous func: functions without name is anonymous func.
// IIFE: if a function can be call/invoked immediately after declaring then it's IIFE.
// function expression: if we assign a function in a variable then it'll be function expression.
// In local scope we can't invoke or call a function literal (function expression) before its assignment. We must invoke it from a lower position, after the function expression is defined.
// parameter: parameter is the variable defined inside a func.
// argument: argument is the value we pass in func by calling or invoking it.
//first order function and higher order function came from functional paradigm, some functional paradigm language is (haskell, racket). These languages are inspired from math's logic especially discrete mathematics, where there are 2 kinds of logic behind first order function and higher order function, 1 first order logic, 2 higher order logic.
// first order func: those who works with things like number, boolean, string are first order func.
// higher order func: if a function receives a function as a parameter or return it or does both then it's a higher order function.
// callback function:  the function we pass through higher order functions are callback function

/*
for a function to be higher order function it must have one of the following rules in them:
	1. take function as parameter
	2. function return
	3. both
*/

/*
1. First order function
	i. standard or named function
	ii. anonymous function
	iii. IIFE
	iv. function expression

2. Higher order function or first class function
3. callback function
4. first call citizen => the data we assign in a variable are the first class citizen


functional paradigm -> haskell, racket

math -> logic (discrete mathematics)

1. first order logic
2. higher order logic

### logic ###

1. Object (people, animal, car etc)
2. Property (color, student)
3. Relation

Tom is a student.
Apple is red.
Tom is taller than Jerry.

Statement

## First order logic:

Rule: All customers must pay their pizza bills.
			All students must wear their uniforms.

## Higher order logic:
			Any rule that applies to all customers must also apply to Tom

			a rule: no customers should give tips for the service

*/
