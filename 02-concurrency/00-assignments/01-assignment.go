package main

import "fmt"

func main() {
	for i := 200; i <= 1000; i++ {
		printIfPrime(i)
	}
}

func printIfPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Printf("Prime : %d\n", no)
}
