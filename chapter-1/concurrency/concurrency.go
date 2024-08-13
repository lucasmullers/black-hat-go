package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("f function")
}

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	// To execute code concurrently, you can use goroutines, which are functions or methods that can run simultaneously
	// Goroutines are called lightweight threads because the cost of creating them is minimal when compared to actual threads
	// To create a goroutine, use the go keyword before the call to a method or function you wish to run concurrently
	fmt.Println("Goroutine")
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	// Go contains a data type called channels that provide a mechanism through which goroutines can synchronize their
	// execution and communicate with each other.
	fmt.Println("\nChannels")
	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)

	// <- operator indicates wheter the data is flowing to or from channel.
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
