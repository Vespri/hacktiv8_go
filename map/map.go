package main

import "fmt"

func main() {
	var person map[string]string //Deklarasi

	person = map[string]string{} //Inisialisasi

	person["name"] = "Kresna"

	person["age"] = "21"

	person["city"] = "Palembang"

	fmt.Println("Name :", person["name"])
	fmt.Println("Age :", person["age"])
	fmt.Println("City :", person["city"])

	// Looping with map
	var people = map[string]string{
		"name": "Kresna",
		"age":  "21",
		"city": "Palembang",
	}

	for key, value := range people {
		fmt.Println(key, ":", value)
	}

	// Deleting value
	var secondPerson = map[string]string{
		"name":    "Vespri",
		"age":     "21",
		"country": "Indonesia",
	}

	fmt.Println("Before deleting :", secondPerson)

	delete(secondPerson, "age")

	fmt.Println("After deleting :", secondPerson)

	// Detecting a value

	var thirdPerson = map[string]string{
		"name":    "Wijaya",
		"age":     "21",
		"address": "Jalan Sekip",
	}

	value, exist := thirdPerson["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(thirdPerson, "age")

	value, exist = thirdPerson["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	// Combining slice with map
	var fourthPerson = []map[string]string{
		{"name": "Kresna", "age": "21"},
		{"name": "Vespri", "age": "21"},
		{"name": "Wijaya", "age": "21"},
	}

	for i, newPerson := range fourthPerson {
		fmt.Printf("Index : %d, name : %s, age : %s\n", i, newPerson["name"], newPerson["age"])
	}
}
