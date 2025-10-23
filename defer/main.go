package main

import "fmt"

// func a() {
// 	i := 0
// 	fmt.Println("first", i)
// 	defer fmt.Println("second", i)
// 	i = i + 1               // or i++
// 	fmt.Println("third", i) // 1
// 	defer fmt.Println("fourth", i)
// 	// return
// }

// func sum(a int, b int) (s int) { //named return value
// 	s = a + b
// 	return
// }

func calculate() (result int) {
	fmt.Println("first", result)

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}

	defer show()

	result = 5
	fmt.Println("second", result)
	// now result =  5, and this 5 will replace show funcs result which is 0 but after replace it'll be 5, and then defer will be 15, and bc results reference holds total sum of result so in the end result = will be 15 and that will be returned so main first = will be 15. and about storing reference? run time managed memory will store the defer and closure.  i need more clarification about this one, there still so many confusion which yet to understand.
	return
}

func calc() int {
	result := 0
	fmt.Println("first", result) // 0

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}
	defer show() //

	result = 5
	fmt.Println("second", result)
	//show yet to execute so from here the value of result will be changed to 5, and since result is a part of closure so it'll be used as referenced and that's how results 0 will be replaced with 5 so show = 15 now. but how it returns
	return result
}

func main() {
	a := calculate()
	b := calc()

	fmt.Println("main first", a)
	fmt.Println("main second", b)
}

/*
defer stays in runtime-managed memory.
Runtime-managed memory = memory that’s allocated, tracked, and freed by the Go runtime, not directly by the OS.

Defer follows LIFO (Last In, First Out) — the last deferred func will be the first one to execute.

defer uses a linked list data structure.
The linked list of defers is stored in memory managed by the Go runtime.

all the defer func forms closure, it's a must.

what are the rules of named return value func and simple return value func?


*/

/*
	after compilation phase a and main func will be in code segment, and when it start to execute there will create main stack frame inside stack, then it'll check wts inside main then it'll get a() so it'll go to check wts inside a(), it'll see i = 0, then println to print the value of i so it'll print it, then it'll see defer func and defer will go in runtime memory to execute later, then it'll see in next line i = i + 1 and next line will print the new value of i, then there again defer which will be taken in runtime managed memory to execute later, now that everything from a() is executed except defer so now runtime managed memory will create a println stack frame to execute defer and it'll follow lifo to execute all the defer is has from a()
*/

/*
after compilation phase, a() and main() will be in the code segment.
when execution starts, a main stack frame is created in the stack.
it checks main, finds a(), and goes inside a().
it sees i = 0, then println prints it.

then it sees defer, which gets registered in runtime-managed memory to run later.
then i = i + 1, and println prints the new value.
then another defer, which also gets registered to run later.

now everything in a() is done except the defers,
so before returning, runtime executes them in LIFO order —
last defer runs first, then the first one.

*/
