/*
In the given solution, one worker is created per number.
Modify the program so that number or workers are fixed (ex:10) and they share the load of processing the numbers
*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	primeNoCh := genPrimes(1000, 2000)
	for primeNo := range primeNoCh {
		fmt.Println("Prime :", primeNo)
	}
}

func genPrimes(start, end int) <-chan int {
	primeNoCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go isPrime(no, primeNoCh, wg)
		}
		wg.Wait()
		close(primeNoCh)
	}()
	return primeNoCh
}

func isPrime(no int, outCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	outCh <- no
}
