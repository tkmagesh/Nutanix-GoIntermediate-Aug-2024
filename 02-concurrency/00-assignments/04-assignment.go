// modify the program so that it follows "share memory by communicating (channels)"

package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory
var primes []int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 200; i <= 1000; i++ {
		wg.Add(1)
		go checkPrime(i, wg)
	}
	wg.Wait()
	for _, primeNo := range primes {
		fmt.Println("Prime :", primeNo)
	}
}

func checkPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	mutex.Lock()
	{
		primes = append(primes, no)
	}
	mutex.Unlock()
}
