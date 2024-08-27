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

	// context to carry the data
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")

	// context for timeout
	timeoutCtx, cancel := context.WithTimeout(valCtx, 10*time.Second)

	// programmatic cancellation
	go func() {
		fmt.Println("Hit ENTER to stop (before 10 secs)...")
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go doWork(timeoutCtx, wg)
	wg.Wait()
	switch {
	case timeoutCtx.Err() == context.Canceled:
		fmt.Println("context cancelled programmatically")
	case timeoutCtx.Err() == context.DeadlineExceeded:
		fmt.Println("context auto cancelled by timeout")
	}

}

// keep doing something until a cancel signal is received
func doWork(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("ctx.Value('root-key') = %v\n", ctx.Value("root-key"))
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
