package main

import (
	"fmt"
)

func main() {
	go f1() // scheduling f1() execution through the scheduler
	f2()
	// app shutdown when the main() is done (irrespective of any goroutines scheduled and waiting for their turn)

	// poor man's synchronization
	// time.Sleep(500 * time.Millisecond)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
