package main

import "fmt"

var arr2 = [3]string{"Tom", "Jerry", "Butch"}

func main() {
	// var arr [2]int
	// arr[0] = 5
	// arr[1] = 10

	arr := [2]int{2, 4} //shorthand

	fmt.Println(arr)
	fmt.Println(arr2)
}

/*
Array:
- An array is a data structure used to store a fixed-size sequence of elements of the same type.
- In Go, array size is part of its type (e.g., [3]int and [4]int are different types).
- Arrays are value types, meaning assignment or passing to a function creates a copy.
- Default value: arrays are initialized with the zero value of their type (0 for int, "" for string, false for bool, etc.).
- Indexing: array indices start at 0.
- Length: use len(arr) to get array size.
- Iteration: arrays can be iterated using for or for range.
- Multidimensional: arrays can contain other arrays, e.g., [2][2]int.
*/
