package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	wg := &sync.WaitGroup{}
	// programmatic cancellation
	cancelCtx, cancel := context.WithCancel(rootCtx)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go doWork(cancelCtx, wg)
	wg.Wait()

}

// keep doing something until a cancel signal is received
func doWork(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[doWork] cancellation signal received")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("doing the work!")
		}
	}
}
