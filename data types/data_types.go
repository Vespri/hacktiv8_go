package main

import "fmt"

func main() {
	// INT
	var first = 1
	var second int8 = 2

	// Float
	var decimalNumber float64 = 3.14

	// String
	var message = "Hai Hacktiv8 team, perkenalkan nama saya Kresna Vespri Wijaya"

	fmt.Printf("Tipe data first : %T\n", first)
	fmt.Printf("Tipe data second : %T\n", second)
	fmt.Printf("decimal Number : %f\n", decimalNumber)
	fmt.Printf("decimal : %.4f\n", decimalNumber)
	fmt.Println(message)
}
