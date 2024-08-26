package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment by 1
	go f1()
	f2()
	wg.Wait() // block until the counter becomes 0 (default)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
