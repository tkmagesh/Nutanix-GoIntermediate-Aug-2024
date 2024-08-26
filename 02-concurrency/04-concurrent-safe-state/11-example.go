package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter atomic.Int64

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count : ", counter.Load())
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Add(1)
}
