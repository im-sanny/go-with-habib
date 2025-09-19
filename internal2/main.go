// internal memory class 2
package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}
	add(1, 1)
	add(p, a)
}

func main() {
	call()
	fmt.Println(a)
}

func init() {
	fmt.Println("I'm init!")
}

/*
	A go program runs in 2 phase:
		1. compilation: here the code turns into binary executables, no code runs here.
		2. execution: compiled binary files start to run here by maintaining execution order.



		************** Compile phase ***************

		** Code segment **
		Code segment: compiled function and constant stay stored here as a part of executable instruction, code segment is read only.
		Data segment: all the global variable stays here.
		Function expression: at runtime this counts as object and stored in code segment.

		a = 10
		call= func(){...}
		add = func(){...}
		main = func(){...}
		init = func(){...}

		go run main.go => compile it => ./main
		go build main.go => compile it => main

		./main
*/

/*
	Visualization of compilation phase ->
	go build main.go

	compiler: checks syntax, scope and dependency.
	saves:
		> constant: a = 10
		> global variable: p = 100
		> function: init, main, call and anonymous function inside(add)
		> generates necessary machine code and binary of metadata.

	What's in the binary file:
		> code segment: a, main, call and anonymous function.
		> data segment: p
		> no execution happens here.

	Visualization of execution phase ->
		1. init()            => "Hello"
   	2. main()            => call() #calls the func
   	3. call()            => add() #declare and call the func
    		- add(5, 6)      => 11
    		- add(100, 10)   => 110
   	4. fmt.Println(a)    => 10

*/
