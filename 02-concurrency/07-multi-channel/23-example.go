/*
Modify the program so that the genNos() keeps producing data until the user hits ENTER key
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genNos(stopCh)
	go func() {
		fmt.Println("Hit ENTER to stop...!")
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done!")
}

func genNos(stopCh chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case ch <- i * 10:
				time.Sleep(500 * time.Millisecond)
			case <-stopCh:
				fmt.Println("stop signal received")
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}
