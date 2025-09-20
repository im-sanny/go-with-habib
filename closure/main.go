package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age =", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("=== Bank ===")
}

/*
Closure: closure is an function who can create access and remember variable outside it's own scope.

A closure is a function that remembers and can access variables from its outer scope, even after the outer function has finished executing.

2 Phases:
	1. compilation phase
	2. execution phase

	*** Code Segment ***

	a = 10
	outer = func(){...}
	call = func(){...}
	main = func(){...}
	init = func(){...}

*/
