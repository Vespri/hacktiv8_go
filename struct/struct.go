package main

import (
	"fmt"
	"strings"
)

// Struct
type Employee struct {
	name     string
	age      int
	division string
}

type Job struct {
	role   string
	person Employee
}

func main() {
	// Giving value to struct
	var employee Employee

	employee.name = "Kresna"
	employee.age = 21
	employee.division = "Software Engineer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
	fmt.Println()

	// Initializing struct
	var employee1 = Employee{}

	employee1.name = "Kresna"
	employee1.age = 21
	employee1.division = "Software Engineer"
	var employee2 = Employee{name: "Vespri", age: 21, division: "Teaching Asistant"}

	fmt.Printf("Employee 1 : %+v\n", employee1)
	fmt.Printf("Employee 2 : %+v\n", employee2)
	fmt.Println()

	// Pointer to a struct
	var employee3 = Employee{name: "Taufiiq", age: 21, division: "Software Engineer"}

	var employee4 *Employee = &employee3

	fmt.Println("Employee 3 name :", employee3.name)
	fmt.Println("Employee 4 name :", employee4.name)

	fmt.Println(strings.Repeat("#", 20))

	employee4.name = "Rian"

	fmt.Println("Employee 3 name :", employee3.name)
	fmt.Println("Employee 4 name :", employee4.name)
	fmt.Println()

	// Embedded struct
	var jobDesk = Job{}

	jobDesk.person.name = "Kresna"
	jobDesk.person.age = 21
	jobDesk.role = "Senior Software Engineer"

	fmt.Printf("%+v", jobDesk)
	fmt.Println()

	// Anonymous struct without filling in fields
	var employee5 = struct {
		people  Employee
		address string
	}{}
	employee5.people = Employee{name: "K", age: 21, division: "CTO"}
	employee5.address = "Palembang"

	// Anonymous struct with filling in fields
	var employee6 = struct {
		people  Employee
		country string
	}{
		people:  Employee{name: "V", age: 21, division: "CEO"},
		country: "Indonesia",
	}

	fmt.Printf("Employee 5 : %+v\n", employee5)
	fmt.Printf("Employee 6 : %+v\n", employee6)

	// Slice of struct
	var employee7 = []Employee{
		{name: "Kresna", age: 21},
		{name: "Vespri", age: 21},
		{name: "Wijaya", age: 21},
	}

	for _, v := range employee7 {
		fmt.Printf("%+v\n", v)
	}

	// Slice of anonymous struct
	var office = []struct {
		employee8 Employee
		address   string
	}{
		{employee8: Employee{name: "Kresna", age: 21}, address: "Palembang"},
		{employee8: Employee{name: "Vespri", age: 21}, address: "Palembang"},
		{employee8: Employee{name: "Wijaya", age: 21}, address: "Palembang"},
	}

	for _, v := range office {
		fmt.Printf("%+v\n", v)
	}
}
