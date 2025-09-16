package main

import "fmt"

// ------------------- Standard or Named Function -------------------
// function with name are standard or named function
func add(a int, b int) { //parameter => a, b
	fmt.Println("Standard Function:", a+b)
}

// ------------------- Higher Order Function -------------------
// if a function receives a function as a parameter or return it or does both
func processOperation(a int, b int, op func(x int, y int)) {
	op(a, b) // op is a callback here
}

// callback function: the function we pass through higher order functions
func addr(x int, y int) {
	z := x + y
	fmt.Println("Callback Function:", z)
}

// higher order example: function returning a function
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// ------------------- init Function -------------------
// u can't call init func, program will call it by it's own,
// init function will be executed before any other function
func init() {
	fmt.Println("Init Function: I'll execute first!")
}

func main() {
	// ---------- standard function ----------
	add(2, 3) //argument => 2, 3

	// ---------- higher order function with callback ----------
	processOperation(5, 6, addr)

	// ---------- anonymous function ----------
	anon := func(a int, b int) {
		sum := a + b
		fmt.Println("Anonymous Function:", sum)
	}
	anon(3, 4)

	// ---------- IIFE (Immediately Invoked Function Expression) ----------
	func(a int, b int) {
		sum := a + b
		fmt.Println("IIFE:", sum)
	}(10, 20)

	// ---------- function expression ----------
	expr := func(a int, b int) {
		c := a + b
		fmt.Println("Function Expression:", c)
	}
	expr(7, 8)

	// ---------- higher order function returning function ----------
	double := multiplier(2) // returns a function that multiplies by 2
	triple := multiplier(3) // returns a function that multiplies by 3
	fmt.Println("Function Return Example:", double(5)) // 10
	fmt.Println("Function Return Example:", triple(5)) // 15
}

/*
Notes:

1. Standard func: function with name are standard or named function
2. init func: u can't call init func, program will call it by it's own, init function will be executed before any other function
3. anonymous func: functions without name is anonymous func.
4. IIFE: if a function can be call/invoked immediately after declaring then it's IIFE.
5. function expression: if we assign a function in a variable then it'll be function expression.
   In local scope we can't invoke or call a function literal (function expression) before its assignment. We must invoke it from a lower position, after the function expression is defined.
6. parameter: parameter is the variable defined inside a func.
7. argument: argument is the value we pass in func by calling or invoking it.

8. first order function and higher order function came from functional paradigm,
   some functional paradigm language is (haskell, racket).
   These languages are inspired from math's logic especially discrete mathematics,
   where there are 2 kinds of logic behind first order function and higher order function,
   1 first order logic, 2 higher order logic.

9. first order func: those who works with things like number, boolean, string are first order func.
10. higher order func: if a function receives a function as a parameter or return it or does both then it's a higher order function.
11. callback function: the function we pass through higher order functions are callback function
12. first class citizen => the data we assign in a variable are the first class citizen

for a function to be higher order function it must have one of the following rules in them:
	1. take function as parameter
	2. function return
	3. both

----------------------------------------------------------

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
