package main

import (
	"fmt"
	"time"
)

func main() {
	go func ()  {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello from main goroutine")
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from main goroutine 2")
	}
	time.Sleep(1 * time.Second) // Wait for goroutine to finish
}
