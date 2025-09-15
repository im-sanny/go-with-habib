package main

import "fmt"

var (
	a = 20
	b = 30
)

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	p := 40
	q := 50

	add(p, q)
	add(a, b)
	add(a, p)
	// add(b, z) //it will give undefined because z is in add func's local scope and local scope can't be access globally.
}
