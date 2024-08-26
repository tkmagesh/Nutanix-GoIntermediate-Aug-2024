package main

import (
	"fmt"
	"sync"
)

func main() {
	run()
}

func run() {
	wg := &sync.WaitGroup{}
	for i := 200; i <= 1000; i++ {
		wg.Add(1)
		printIfPrime(i, wg)
	}
	wg.Wait()
}

func printIfPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Printf("Prime : %d\n", no)
}
