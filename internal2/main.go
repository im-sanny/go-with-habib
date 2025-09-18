// internal memory class 2
package main

import "fmt"

const b = 10

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}
	add(1, 1)
	add(p, b)
}

func main() {
	call()
}

/*
	2 phase
		1. compilation
		2. execution

		************** Compile phase ***************

		** Code segment **
		a = 10
		call= func(){...}
		add = func(){...}
		main = func(){...}
		init = func(){...}

		go run main.go => compile it => ./main
		go build main.go => compile it => main

		./main
*/
