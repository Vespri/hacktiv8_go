package exported_unexported

import "fmt"

func Greet() {
	fmt.Println("Hi from Greet function")
}

func greet() {
	fmt.Println("Hi from greet function")
}
