package main

import "fmt"

var a = 10

func main() {
	age := 30

	if age < 18 {
		a := 47
		fmt.Println(a)
	}
	fmt.Println(a)
}

//variable shadowing happens when a local scope variable name match's with any outer scope variable
