package main

import (
	"fmt"
)

// only signature of the functions
type People interface {
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

type BankUser interface {
	WithdrawMoney(amount float64) float64
}

type user struct {
	Name  string
	Age   int
	Money float64
}

// receiver methods (value receivers → work on copies)
func (obj user) PrintDetails() {
	fmt.Println("Name:", obj.Name)
	fmt.Println("Age:", obj.Age)
	fmt.Println("Money:", obj.Money)
}

func (obj user) WithdrawMoney(amount float64) float64 {
	obj.Money -= amount
	return obj.Money
}

func (obj user) ReceiveMoney(amount float64) float64 {
	obj.Money += amount
	return obj.Money
}

func main() {
	// interface assignment → instantiation
	var usr1 People = user{
		Name:  "U2Fubnk",
		Age:   69,
		Money: 88,
	}

	var usr3 BankUser = user{
		Name:  "Fumis",
		Age:   20,
		Money: 33.99,
	}

	// works on copy (value receiver)
	usr3.WithdrawMoney(3.99)

	usr1.PrintDetails()
}

/*
	# Paradigms
	1. Structured Programming Paradigm
	2. Object-Oriented Programming Paradigm
	3. Functional Programming Paradigm

	Singleton Design Pattern:
	- Only one object (instance) is ever created.
	- That single instance is shared by everyone who needs it.
	- Often used for config, logging, database connection, etc.

	Philosophy:
	- The principles or rules behind how something should be designed or approached.

	Design Pattern:
	- A reusable solution to a common software design problem.
	- Comes from OOP concepts, but Go also uses them even though Go is not fully OOP.

	Abstraction:
	- A way to hide details and show only the necessary idea.
	- In OOP, abstraction means exposing behavior but hiding implementation.

	Interface:
	- Pure abstraction. No details, no fields, no implementation.
	- Example: I’m giving you a book that you can definitely open, read, and close. I’m not telling you the title, the author, how many pages, what it looks like, or what the story is — but I guarantee those three actions will work.
	- Interface guarantees the *behavior*, not the structure.

	More precise:
	- An interface = “Here is a thing that has these methods: open(), read(), close().
	  I don’t care how these methods work or what the object looks like.”

	Instantiation:
	- The act of creating an instance or object from a type or struct.

	Signature:
	- A function's name + its parameters + its return types.
	- Does NOT include the body (actual code inside the function).

	Constructor is a function that creates function.

	explecit, 
*/
