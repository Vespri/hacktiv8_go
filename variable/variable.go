package main

import "fmt"

func main() {
	// First way
	var name string = "Kresna"
	var age int = 21

	// Second way
	var phone, citizen string

	phone, citizen = "08987654321", "Palembang"

	// Third way
	degree, major := "D3", "Full Stack Web Developer"
	dateOnly := 23

	// Forth way
	var satu, dua, tiga = "1", 2, 3

	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("Phone Number : ", phone)
	fmt.Println("Citizen : ", citizen)
	fmt.Println("Degree : ", degree)
	fmt.Println("Major : ", major)
	fmt.Println(satu, dua, tiga)
	fmt.Printf("Hai nama saya %s, umur %d tahun \n", name, age)
	fmt.Printf("%T, %T  ", degree, dateOnly)
}
