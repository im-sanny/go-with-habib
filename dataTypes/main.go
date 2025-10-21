package main

import "fmt"

func main() {
	var a int8 = -128
	var b int8 = 127

	var x uint8 = 10

	var i float32 = 8.477779789798
	var j float64 = 6.9897980

	var flag bool = true

	var s string = "I'm Tom"

	r := '❤'

	fmt.Printf("%c\n", r)   //character -> printf("%c", r)
	fmt.Printf("%d\n", a)   // decimal -> d
	fmt.Printf("%.2f\n", j) // floating -> f -> Printf("%f.3f", j)
	fmt.Printf("%v\n", flag)

	fmt.Printf("%s\n", s)
	fmt.Printf("%T\n", s) // this for getting the type from the variable

	fmt.Println(a, b, x, i, j, flag)
}

/*
int is system-dependent — it adjusts its size based on the system architecture. On 32-bit systems it usually holds 4 bytes, and on some 64-bit systems it can hold 8 bytes.

int8 can only keep value from -128 to 127 also can't hold bigger or smaller than this, it can hold 1 byte.

uint: unsigned int, this type only accepts positive numeric, starts from 0.

float: used to store fractional or decimal numbers, mainly for representing values that aren’t whole numbers.

in a 32 bit computer we can assign float64, and this will take 2 cell of the ram to store the value.

boolean: used to store only two possible values: true or false. each bool value takes 8 bit or 1 byte space.

byte: 8 bits of data, commonly used to represent a character or a small number.
byte -> alias for uint8 -> 8 bits per character -> 1 byte

**rune: a type used to store a single Unicode character.
rune -> alias for int32(unicode point) -> 32 bits -> 4 bytes -> %c

*/
