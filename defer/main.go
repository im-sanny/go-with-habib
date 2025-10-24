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

	// now result = 5, and it holds this new value in memory.
	// later when show() runs, it looks at the same result (by reference),
	// so it sees the updated value (5), adds 10 → result = 15.
	// since result is a named return variable, this updated value (15)
	// becomes the final return value, so main prints 15.
	// about storing reference: the runtime stores the defer and the closure
	// in its managed memory (runtime-managed space), not in the stack.

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
	//now return value will be copied before defer run, bc in this type of func return don't care about what defer will return or what value it has
	return result
}

func main() {
	a := calculate()
	b := calc()

	fmt.Println("main first", a)
	fmt.Println("main second", b)
}

/*
Defer ->

Defer is a Go keyword that delays the execution of a function until the surrounding function returns.

defer stays in runtime-managed memory.
Runtime-managed memory = memory that’s allocated, tracked, and freed by the Go runtime, not directly by the OS.

Defer follows LIFO (Last In, First Out) — the last deferred func will be the first one to execute.

defer uses a linked list data structure.
The linked list of defers is stored in memory managed by the Go runtime.

all the defer func forms closure, it's a must. is it?

calculate():
- named return → result created at start.
- defer stored in runtime memory, closure sees result by reference.
- result = 5 → prints second 5.
- on return → defer runs, adds 10 → result = 15.
- final return = 15.

calc():
- normal return → result is local only.
- return result (5) happens first, then defer runs (makes result 15).
- but return already copied 5 → final return = 5.

- named return → defer can change return.
- normal return → defer runs after value copied.
- defer stored in runtime (linked list), closure captures by reference.

named return: return value is live till the end, so defer can still change it before returning.
normal return: return value gets copied first, then defer runs, so defer can't affect it.


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
