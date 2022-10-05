package main

import "fmt"

func main() {
	var score = 6

	if score > 7 {
		switch {
		case score == 10:
			fmt.Println("Perfect")
		default:
			{
				fmt.Println("Nice!")
			}
		}
	} else {
		if score < 8 {
			fmt.Println("Not bad")
		} else if score == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if score == 0 {
				fmt.Println("Try harder")
			}
		}
	}
}
