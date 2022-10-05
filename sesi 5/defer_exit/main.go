package main

import (
	"fmt"
	"os"
)

func main() {
	// Defer #1
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go Learning Center")

	// Defer #2
	callDeferFunc()
	fmt.Println("Hello There !!")

	// Exit
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func start to execute")
}
