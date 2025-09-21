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
Closure: A closure is a function that can access and remembers variables outside its own scope, even after that scope has finished executing.

2 Phases:
	1. compilation phase(compile time)
	2. execution phase(run time)

	*** Code Segment ***
	a = 10
	outer = func(){...}
	call = func(){...}
	main = func(){...}
	init = func(){...}



	go run main.go => compile it => main => ./main
	go build main.go => compile it = main

	./main

	Escape Analysis: If a variable needs to outlive its own function and needed to use later then Go compiler sends it to heap, this is called escape analysis.

		Escape analysis is the process where the Go compiler determines if a variable must outlive the scope (usually the function) in which it was created. If it does, the variable 'escapes' to the heap; otherwise, it can be efficiently allocated on the stack.

	# stack can be clean automatically but GC cleans heap.


*/
