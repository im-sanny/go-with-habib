package main

import (
	"fmt"

	"example.com/mathlib"
)

var (
	a = 20
	b = 30
)

func main() {
	fmt.Println("Creating Custom Package")
	mathlib.Add(a, b)
	mathlib.Add(8, mathlib.Money)
}
