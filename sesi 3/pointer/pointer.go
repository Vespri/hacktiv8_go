package main

import (
	"fmt"
	"strings"
)

func main() {
	// Pointer
	var firstName *int
	var middleName *int
	_, _ = firstName, middleName

	// Memory address
	var firstNumber int = 2
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber)

	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memori address) :", secondNumber)
	fmt.Println()

	// Changing value through pointer
	var firstPerson string = "Kresna"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondPerson (value) :", *secondPerson)
	fmt.Println("secondPerson (memori address) :", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Vespri"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondPerson (value) :", *secondPerson)
	fmt.Println("secondPerson (memori address) :", secondPerson)
	fmt.Println()

	// Pointer as a parameter
	var a int = 10

	fmt.Println("Before :", a)

	changeValue(&a)

	fmt.Println("After :", a)
}

func changeValue(number *int) {
	*number = 20
}
