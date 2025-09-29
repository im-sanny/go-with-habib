package main

import "fmt"

// Example: Slice from an array
// func sliceFromArray() {
// 	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
// 	s := arr[1:4] // ["is", "a", "Go"]
// 	fmt.Println(s)

// 	s1 := s[1:2] // ["a"]
// 	// capacity is 5-1 = 4, because slice uses underlying array reference
// 	fmt.Println(s1)
// 	fmt.Println("len:", len(s1), "cap:", cap(s1))
// }

// Example: Change slice inside a function
// func changeSlice(a []int) []int {
// 	a[0] = 10         // modifies underlying array
// 	a = append(a, 11) // may reallocate if capacity is full
// 	return a
// }

func main() {
	// Slice literal
	// s1 := []int{1, 2, 3}
	// fmt.Println("slice:", s1, "len:", len(s1), "cap:", cap(s1))

	// make with len only
	// s2 := make([]int, 3) // [0 0 0], len=3, cap=3
	// s2[0] = 69
	// fmt.Println(s2, "len:", len(s2), "cap:", cap(s2))

	// make with len and cap
	// s3 := make([]int, 3, 5) // [0 0 0], len=3, cap=5
	// s3[2] = 10
	// fmt.Println(s3, "len:", len(s3), "cap:", cap(s3))

	// empty slice (nil slice)
	// var s4 []int
	// s4 = append(s4, 1, 2, 3)
	// fmt.Println(s4, "len:", len(s4), "cap:", cap(s4))

	// append capacity growth
	var x []int
	x = append(x, 1) // [1], len=1, cap=4 (Go reserves space upfront)
	x = append(x, 2) // [1 2], len=2, cap=4
	x = append(x, 3)

	y := x

	x = append(x, 4)
	y = append(y, 5) // here y doesn't have it's own array it's share the same array with x, so if y tries to change or add anything then it'll directly change x, also y and x share same pointer, before this point y had len 3 and cap 4 so there was one empty place before coming to this point, and when y append 5 it goes to that empty place y had and when it sits on the empty place y array changed and since y array is x under the hood it will overwrite 4 and put 5

	x[0] = 10

	fmt.Println(x)
	fmt.Println(y)


	// Slice sharing and append
	// x = []int{1, 2, 3, 4, 5}
	// x = append(x, 6, 7)
	// a := x[4:]          // slice of last elements [5 6 7]
	// y := changeSlice(a) // modifies slice
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
}

/*
---------------------------------------
NOTES ON SLICES
---------------------------------------

1. Slice is not an array.
   Slice is built on top of an array and is more flexible.

2. Internally a slice has 3 parts:
   - pointer: points to the underlying array
   - length: number of current elements
   - capacity: maximum elements it can grow before reallocation

3. Ways to create a slice:
   1. From an array → arr[1:4]
   2. From another slice → s[1:2]
   3. Slice literal → []int{1,2,3}
   4. Using make(len) → make([]int, 3)
   5. Using make(len, cap) → make([]int, 3, 5)
   6. Nil slice → var s []int
   7. Append function → append(s, 1,2,3)

4. Append function:
   - Expands the length by adding new elements
   - If no capacity left, it reallocates a bigger array
   - After Go 1.18 update:
       * At first append, capacity becomes 4 (not 1,2,4 like before)
       * Until 256 → doubles capacity
       * 256–512 → increases by ~1.25x
       * 512–1024 → increases by ~1.5x
       * After 1024 → increases by 25%

5. Array vs Slice:
   - Array has fixed size, cannot grow
   - Slice can grow dynamically with append

6. Index rule:
   - Array of size 10 → max index = 9
   - If you try arr[10] = 55 → runtime error (index out of range)

7. Variadic functions also use slices internally.
*/
