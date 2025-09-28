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

	// var s []int            //empty slice or nil slice //[], len = 0, cap = 0
	// s = append(s, 1, 2, 3) //[, 2, 3], len = 3, cap = 3
	// fmt.Println(s)

	var x []int      //[], len = 0, cap= 0
	x = append(x, 1) //[1], len = 1, cap = 1
	x = append(x, 2) //[1, 2] len = 2, cap = 2
	x = append(x, 3) //[1, 2, 3] len = 3, cap = 4 // cap become 4 bc when theres no space left and there sth needs to add then it follows slice underlying rules and by that the cap expands.
	fmt.Println(x)   //[1, 2, 3]

	y := x

	x = append(x, 4) //[]
	y = append(y, 5)
	x[0] = 10

	fmt.Println(x)
	fmt.Println(y)

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
	8. slice underlying array rules => 1024 -> 100% increase, until 1024 it will increase by 100%, after 1024 it will increase by 25%.
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

/*
	x = append(x, 1)
	x = append(x, 2)
	here append will expand the length and capacity of the array each time we append new value in it. and this is how we can expand length and capacity of the array without any issue.
*/
