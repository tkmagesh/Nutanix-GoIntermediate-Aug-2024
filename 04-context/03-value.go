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
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")
	wg.Add(1)
	go doWork(valCtx, wg)
	wg.Wait()

}

// keep doing something until a cancel signal is received
func doWork(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("ctx.Value('root-key') = %v\n", ctx.Value("root-key"))
	for range 10 {
		time.Sleep(1 * time.Second)
		fmt.Println("doing the work!")

	}
}
