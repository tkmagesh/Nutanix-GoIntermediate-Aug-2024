package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(<-ch3)
	}()

	// order of consumption is pre-determined
	/*
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	*/

	// order of consumption is NOT pre-determined
	// select-case == switch case for channels
	for range 3 {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case d2 := <-ch2:
			fmt.Println(d2)
		case ch3 <- 300:
			fmt.Println("[select-case] data sent to ch3")
		}
	}
}
