package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for no := range ch {
		fmt.Println(no)
	}
}

func genNos() <-chan int {
	ch := make(chan int)
	timeoutCh := time.After(5 * time.Second)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case ch <- i * 10:
				time.Sleep(500 * time.Millisecond)
			case <-timeoutCh:
				fmt.Println("timeout occurred!")
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}

/*
func timeout(d time.Duration) <-chan time.Time {
	timeOutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeOutCh <- time.Now()
	}()
	return timeOutCh
}
*/
