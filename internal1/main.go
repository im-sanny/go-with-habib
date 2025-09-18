// internal memory class 1
package main

import "fmt"

var a = 10

func add(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	add(1, 2)
	add(a, 5)

}

func init() {
	fmt.Println("I'm init!")
}

/*
	Code segment: all the functions goes to the code segment after compile. it stores things that has no chance to change later like: functions, const variables.
	Data segment: everything in global memory goes to the data segment or data segment keeps everything global.
	Stack: local variable used inside a function stored in stack, and for every function call there creates a new stack frame.
	The stack is a memory area that temporarily stores data for currently executing functions, including their local variables and call information.
	Heap:
	GC: it manages heap

	stack frame: in stack a function takes some space and that space is called that functions stack frame
	compiler reads all the code at first before executing
*/
