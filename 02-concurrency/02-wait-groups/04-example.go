package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1) // increment by 1
		go f1(wg)
	}
	f2()
	wg.Wait() // block until the counter becomes 0 (default)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement by 1
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
