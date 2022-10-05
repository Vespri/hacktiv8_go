package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel operators
	c := make(chan string)

	// c <- value

	// result := <-c

	// Implementing channel
	go introduce("Kresna", c)

	go introduce("Vespri", c)

	go introduce("Wijaya", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	fmt.Println()

	// Channel with anonymous function
	students := []string{"Norman", "Rian", "Taufiiq"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hi, my name is %s\n", student)
			c <- result
		}(v)
	}

	for i := 1; i < 4; i++ {
		print(c)
	}

	close(c)

	// Buffered vs unbuffered channel
	k := make(chan int)

	go func(v chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		v <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(k)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	w := <-k
	fmt.Println("main goroutine received data:", w)

	close(k)
	time.Sleep(time.Second)

	fmt.Println()

	// Select
	kv := make(chan string)
	vw := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		kv <- "Hello There !!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		vw <- "Hi There !!"
	}()

	for i := 1; i < 3; i++ {
		select {
		case message1 := <-kv:
			fmt.Println("Received", message1)
		case message2 := <-vw:
			fmt.Println("Received", message2)
		}
	}
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hi, my name is %s", student)

	c <- result
}

func print(c chan string) {
	fmt.Println(<-c)
}
