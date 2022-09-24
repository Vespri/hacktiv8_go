package main

import "fmt"

func main() {
	// First way of looping
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	// Second way of looping
	var j = 0
	for j < 3 {
		fmt.Println("Angka", j)
		j++
	}

	// Third way of looping
	var k = 0
	for {
		fmt.Println("Angka", k)

		k++
		if k == 3 {
			break
		}
	}

	// Break and continue keywords
	for l := 1; l < 10; l++ {
		if l%2 == 1 {
			continue
		}

		if l > 8 {
			break
		}

		fmt.Println("Angka", l)
	}

	// Nested Looping
	for m := 0; m < 5; m++ {
		for n := m; n < 5; n++ {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}

	// Label
outerLoop:
	for o := 0; o < 3; o++ {
		fmt.Println("Perulangan ke -", o+1)
		for p := 0; p < 3; p++ {
			if o == 2 {
				break outerLoop
			}
			fmt.Print(p, " ")
		}
		fmt.Print("\n")
	}
}
