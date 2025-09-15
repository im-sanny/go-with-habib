package mathlib

import "fmt"

var Money = 100

func Add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

// to access func variable and types of this or any package from other package we have to name the func with uppercase letter, neither we can't access the file from any other file.
// to run and import mathlib package we need to run the command: go mod init example.com, instead of example.com we can just put any name or domain but it's recommended to use sth unique to maintain standard.
