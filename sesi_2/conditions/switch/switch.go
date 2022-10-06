package main

import "fmt"

func main() {
	// Switch with relational operators
	var score = 6

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Not bad")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You need to learn more")
		}
	}

	// Switch (fallthrough keyword)
	var score2 = 6

	switch {
	case score2 == 8:
		fmt.Println("Perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("Not bad")
		fallthrough
	case score2 < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You need to learn more")
		}
	}
}
