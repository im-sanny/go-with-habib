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

	fmt.Printf("%c\n", r)   // character -> %c
	fmt.Printf("%d\n", a)   // decimal -> %d
	fmt.Printf("%.2f\n", j) // float -> %.2f
	fmt.Printf("%v\n", flag)
	fmt.Printf("%s\n", s) // string -> %s
	fmt.Printf("%T\n", s) // get the type of variable -> %T

	fmt.Println(a, b, x, i, j, flag)
}

/*
Data types ->

int:
    System-dependent — adjusts its size based on the system architecture.
    On 32-bit systems it usually holds 4 bytes,
    and on some 64-bit systems it can hold 8 bytes.

int8:
    Can only keep values from -128 to 127.
    It can hold 1 byte.

uint:
    Unsigned integer — only accepts positive numbers (starts from 0).
    It’s system-dependent and doesn’t have an exact fixed size.

float:
    Used to store fractional or decimal numbers (non-whole values).
    Go doesn’t have a plain “float” type — the default is float64.
    float32 → around 6–7 digits of precision.
    float64 → around 15–16 digits of precision.
    On 32-bit computers, float64 still works — it just takes 2 memory cells.

boolean:
    Used to store only two possible values: true or false.
    Each bool value takes 8 bits (1 byte) of space.

byte:
    8 bits of data, commonly used to represent a character or small number.
    byte → alias for uint8 → 8 bits → 1 byte.

rune:
    Used to store a single Unicode character.
    rune → alias for int32 (Unicode code point) → 32 bits → 4 bytes → %c.

string:
    Sequence of bytes that represent text, stored as UTF-8 encoded data.
    Each character may take 1 or more bytes depending on the symbol.

Format Specifiers:
    %v → prints value
    %T → prints type
    %c → prints character
    %f → prints float
    %d → prints decimal (integer)
    %s → prints string


Go runtime automatically handles how each variable type is stored.
*/
