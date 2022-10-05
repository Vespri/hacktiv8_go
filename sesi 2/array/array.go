package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var names = [3]string{"Kresna", "Vespri", "Wijaya"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", names)

	// Modify Element Through Index
	var fruits = [3]string{"Semangka", "Anggur", "Mangga"}
	fruits[0] = "Srikaya"
	fruits[1] = "Strawberry"
	fruits[2] = "Mangga"

	fmt.Printf("%#v\n", fruits)

	// Loop through elements
	var animals = [3]string{"Kucing", "Semut", "Laba-laba"}
	// First way
	for i, v := range animals {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))
	// Second way
	for i := 0; i < len(animals); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, animals[i])
	}

	// Multidimensional Array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
