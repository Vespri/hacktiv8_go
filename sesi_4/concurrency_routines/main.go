package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go goroutine()

	// Asynchronous process
	fmt.Println("main execution started")

	go firstProccess(8)
	secondProccess(8)

	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("main execution ended")
}

func goroutine() {
	fmt.Println("Hi!")
}

func firstProccess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProccess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}
