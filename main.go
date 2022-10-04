package main

import (
	"fmt"
	"hacktiv8/exported_unexported"
)

func main() {
	fmt.Println("Hello World")

	exported_unexported.Greet()

	var person = exported_unexported.Person{}

	person.Invokegreet()
}
