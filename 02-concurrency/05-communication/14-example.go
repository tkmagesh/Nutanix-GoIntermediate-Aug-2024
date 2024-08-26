package main

import (
	"fmt"
	"sync"
)

/*
func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)
}
*/

// modify the below in such a way that the "send" happens in main() but the "receive" and print happens in a different goroutine
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100
	wg.Wait()

}
