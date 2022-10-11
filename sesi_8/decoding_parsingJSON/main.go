package main

import (
	"encoding/json"
	"fmt"
)

// Decoding JSON To Struct
type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	// Decoding JSON To Struct
	var jsonString = `
	{
		"name": "Kresna",
		"email": "kresna@gmail.com",
		"age": 21
	}
	`

	var result Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Full Name :", result.Name)
	fmt.Println("Email :", result.Email)
	fmt.Println("Age :", result.Age)
}
