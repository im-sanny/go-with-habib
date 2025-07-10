package main

import "fmt"

//passed by value(normal way)
//passed by reference(pointer way)

type User struct {
	Name     string
	Age      int
	Employed bool
}

func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func main() {
	//pointer is a address
	//pointer or address of memory(RAM)

	san := User{ //instance or object
		Name:     "Sanny",
		Age:      25,
		Employed: false,
	}

	p := &san
	fmt.Println(*&p.Name)
	fmt.Println(p.Name) //I'll get same output by using this

	x := 10

	addr := &x  //address of x, ampersand(&) refers address of x here in this line
	*addr = 100 // assigning new value of x through pointer,

	fmt.Println("Address of x:", addr) // addr is the address of x
	fmt.Println("Value of x:", *addr)  // by using * we can get the value of addr, here * used to get the value of the address

	arr := [3]int{1, 2, 3}

	print(&arr)
}

// ampersand & => address of
// * => value at address
