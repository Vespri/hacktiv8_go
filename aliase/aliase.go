package main

import "fmt"

func main() {
	// Declarating a variable with uint8 data type
	var a uint8 = 10
	var b byte // byte is alias from uint8 data type

	b = a // no error, because byte has same data type with uint8
	_ = b

	// Declarate data type names second
	type second = uint

	var hour second = 3600
	fmt.Printf("Hour type : %T\n", hour) // => hour type = uint
}
