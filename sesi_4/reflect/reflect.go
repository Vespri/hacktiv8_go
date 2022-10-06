package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func main() {
	// Identifying Data Type & Value
	var number = 21
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Variable Type :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Variable Value :", reflectValue.Int())
	}

	// Accessing Value Using Interface Method
	var value = "Kresna"
	var reflectValue1 = reflect.ValueOf(value)

	fmt.Println("Variable Type :", reflectValue1.Type())
	fmt.Println("Variable Value :", reflectValue1.Interface())

	// Identifying Method Information
	var s1 = &student{Name: "Kresna Vespri", Grade: 2}
	fmt.Println("Nama :", s1.Name)

	var reflectValue2 = reflect.ValueOf(s1)
	var method = reflectValue2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Vespri"),
	})

	fmt.Println("Nama :", s1.Name)
}

// Identifying Method Information
func (s *student) SetName(name string) {
	s.Name = name
}
