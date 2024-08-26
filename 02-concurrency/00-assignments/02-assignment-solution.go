// Print the total number of prime numbers found in the main() function

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func main() {
	wg := &sync.WaitGroup{}
	for i := 200; i <= 1000; i++ {
		wg.Add(1)
		printIfPrime(i, wg)
	}
	wg.Wait()
	fmt.Println("Count :", count)
}

func printIfPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	atomic.AddInt64(&count, 1)
	fmt.Printf("Prime : %d\n", no)
}
