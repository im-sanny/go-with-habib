package main

import "fmt"

// func slice() {
// 	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}

// 	s := arr[1:4] //[is a go]
// 	fmt.Println(s)

// 	s1 := s[1:2] //[a] capacity is 4 bc it's using main array reference and if we count from a then we'll get the capacity.
// 	fmt.Println(s1)
// 	fmt.Println(len(s1))
// 	fmt.Println(cap(s1))
// }

func main() {
	// s := []int{1, 2, 3} // if i remove fixed size from an array then it becomes slice, and now this calls slice literal
	// fmt.Println("slice: ", s, "length: ", len(s), "capacity: ", cap(s))

	// s := make([]int, 3) //[0, 0, 0] len = 3, cap =3
	// s[0] = 69           //[5, 0, 0] len = 3, cap =3

	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// s := make([]int, 3, 5) //[0, 0, 0] len = 3, cap = 5

	// s[0] = 69 //[5, 0, 0] len = 3, cap =5
	// s[2] = 10
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	var s []int            //empty slice or nil slice //[], len = 0, cap = 0
	s = append(s, 1, 2, 3) //[, 2, 3], len = 3, cap = 3
	fmt.Println(s)

}

/*
Kind of Slices:
	1. slice from an array
	2. slice from a slice
	3. slice literal
	4. make func with len
	5. make func with len and capacity
	6. empty slice or nil slice
	7. add value with append func
*/

/*
	In go a slice has 3 part:
		1. pointer = here "is" is the pointer
		2. length = length is 3
		3. capacity = capacity is 5

		# max index of a 10 size array is 9 bc for array count starts with 0.
		# in a 10 index array u can't define value s[10] = 55, if we try to do this then we'll get runtime error(index out of range).
		# with append func we can add value in empty or in slice that already have value in it.
*/
