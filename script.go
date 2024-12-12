package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Println("Hello, ", name)
}

func main() {
	go greet("Alice") // start a goroutine
	go greet("Bob")   // start another goroutine

	// Wait for goroutines to complete (simple way, not the best practice)
	time.Sleep(time.Second)
	fmt.Println("Main function exits")
}
