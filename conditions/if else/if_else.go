package main

import "fmt"

func main() {
	var currentYear = 2021

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Umur kamu", age)
		fmt.Println("Kamu belum boleh membuat kartu SIM")
	} else {
		fmt.Println("Umur kamu", age)
		fmt.Println("Kamu sudah boleh membuat kartu SIM")
	}
}
