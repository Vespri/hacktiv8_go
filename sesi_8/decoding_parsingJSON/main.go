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

	// Decoding JSON To Map
	var mapResult map[string]interface{}

	var errs = json.Unmarshal([]byte(jsonString), &mapResult)

	if errs != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Full Name :", mapResult["name"])
	fmt.Println("Email :", mapResult["email"])
	fmt.Println("Age :", mapResult["age"])

	// Decoding JSON To Empty Interface
	var temp interface{}

	var err2 = json.Unmarshal([]byte(jsonString), &temp)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}

	var interfaceResult = temp.(map[string]interface{})

	fmt.Println("Full Name :", interfaceResult["name"])
	fmt.Println("Email :", interfaceResult["email"])
	fmt.Println("Age :", interfaceResult["age"])
}
