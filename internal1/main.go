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
	1. Code segment: all the functions goes to the code segment after compile. it stores things that has no chance to change later like: functions, const variables.

	2. Data segment: everything in global memory goes to the data segment or data segment keeps everything that is global.

	3. Stack: local variable used inside a function stored in stack, and for every function call there creates a new stack frame.
		A. The stack is a memory area that temporarily stores data for currently executing functions, including their local variables and call information.
		B. when any function executes, it goes to stack and creates stack frame, those variable that stay in stack frame they also stays in stack.

	4. Stack frame: in stack a function takes some space and that space is called that functions stack frame

	5. Heap: used for dynamic memory allocation.

	6. GC: it manages heap, automatically removes unused or unnecessary memory from heap

	compiler reads all the code at first before executing
	Yes.

	When a stack creates?
		Ans: The call stack is created by the operating system when the program starts executing. A new stack frame is created on that stack each time a function is called during execution.
*/
