package main

import "fmt"

func main(){
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}

	s := arr[1:4]
	fmt.Println(s)
}


/*
	In go a slice has 3 part:
		1. pointer = here "is" is the pointer
		2. length = length is 3
		3. capacity = capacity is 5
*/
