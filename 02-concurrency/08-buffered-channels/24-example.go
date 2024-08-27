package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 10
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 20
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 30
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
}
