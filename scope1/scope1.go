package main

import "fmt"

var (
	c = 10
	d = 20
)

func sum(x int, y int) {
	res := x + y
	printNum(res)
}

func main() {
	sum(c, d)
}

func printNum(num int) {
	fmt.Println(num)
}

//in this example it doesn't matter the position of the func and wether it's align properly or not, it'll work even if u keep the func anywhere of the file but the func must have no error to execute, because go compiler doesn't care about the order of the func declaration.
