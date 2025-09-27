package main

import "fmt"

func slice(){
		arr := [6]string{"This", "is", "a", "Go", "interview", "question"}

	s := arr[1:4]	//[is a go]
	fmt.Println(s)

	s1 := s[1:2] //[a] capacity is 4 bc it's using main array reference and if we count from a then we'll get the capacity.
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

func main(){
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
