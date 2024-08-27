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

	primeNoCh := genPrimes(1000, 2000, 10)
	for primeNo := range primeNoCh {
		fmt.Println("Prime :", primeNo)
	}
}

func genPrimes(start, end int, workerCount int) <-chan int {
	primeNoCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		inputCh := make(chan int)
		go func() {
			for no := start; no <= end; no++ {
				inputCh <- no
			}
			close(inputCh)
		}()
		for range workerCount {
			wg.Add(1)
			go isPrime(inputCh, primeNoCh, wg)
		}
		wg.Wait()
		close(primeNoCh)
	}()
	return primeNoCh
}

func isPrime(inputCh <-chan int, outCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := range inputCh {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		outCh <- no
	}
}
