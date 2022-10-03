package main

import (
	"fmt"
	"strings"
)

type isOddNum func(int) bool

func main() {
	// Declare closure in variable
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	fmt.Println(evenNumbers(numbers...))

	// IIFE
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)

	// Closure as a return value
	var studentLists = []string{"Kresna", "Vespri", "Wijaya"}

	var find = findStudent(studentLists)

	fmt.Println(find("vw"))

	// Callback
	var findOddNumber = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers :", findOddNumber)
}

// Closure as a return value
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s doesn't exist!!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position+1)
	}
}

// Callback
func findOddNumbers(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
