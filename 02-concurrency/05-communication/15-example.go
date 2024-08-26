package main

import (
	"fmt"
)

// consumer
func main() {
	ch := make(chan int)
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	fmt.Println(<-ch)
}

// producer
func add(x, y int) int {
	return x + y
}
