package main

import "fmt"

func main() {
	// Arithmetic Operators
	var valueFirst = 2 + 2*3
	var valueSecond = (2 + 2) * 3

	fmt.Println("2 + 2 * 3 = ", valueFirst)
	fmt.Println("(2 + 2) * 3 = ", valueSecond, "\n")

	// Relational Operators
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first Condition: ", firstCondition)
	fmt.Println("second Condition: ", secondCondition)
	fmt.Println("third Condition: ", thirdCondition)
	fmt.Println("fourth Condition: ", fourthCondition, "\n")

	// Logical Operators
	var correct = true
	var wrong = false

	var correctANDwrong = correct && wrong
	fmt.Printf("Correct && Wrong \t(%t) \n", correctANDwrong)

	var correctORwrong = correct || wrong
	fmt.Printf("Correct || Wrong \t(%t) \n", correctORwrong)

	var correctReverse = !correct
	fmt.Printf("!Correct \t\t(%t) \n", correctReverse)

}
