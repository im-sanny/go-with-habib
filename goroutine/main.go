package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func printHello(num int) {
	fmt.Println("hello fumis", num)
}

func main() {
	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)

	fmt.Println(a, "", p)

	time.Sleep(5 * time.Second)
}

/*
Goroutine:
- Go's concurrent lightweight execution unit
- go runtime is a virtual operating system
- when a go program runs that time it runs it's own mini os which calls go runtime

- lightweight thread/virtual thread.
- works like logical thread.
- executes many function concurrently.
- go runtime manages it.

- before any function if u put the keyword go then it becomes goroutine.

*/
