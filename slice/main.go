package main

import "fmt"

func slice() {
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}

	s := arr[1:4] //[is a go]
	fmt.Println(s)

	s1 := s[1:2] //[a] capacity is 4 bc it's using main array reference and if we count from a then we'll get the capacity.
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

func main() {
	// s := []int{1, 2, 3} // if i remove fixed size from an array then it becomes slice, and now this calls slice literal
	// fmt.Println("slice: ", s, "length: ", len(s), "capacity: ", cap(s))

	s := make([]int, 3) //[0, 0, 0] len = 3, cap =3
	s[0] = 69           //[5, 0, 0] len = 3, cap =3
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	slice()
}

/*
	1. slice from an array
	2. slice from a slice
*/

/*
	In go a slice has 3 part:
		1. pointer = here "is" is the pointer
		2. length = length is 3
		3. capacity = capacity is 5
*/
