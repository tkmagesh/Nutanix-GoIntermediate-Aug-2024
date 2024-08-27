package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	// context root
	rootCtx := context.Background()

	// context to share data
	rootValCtx := context.WithValue(rootCtx, "root-key", "root-value")

	// context to send cancel signal
	cancelCtx, cancel := context.WithCancel(rootValCtx)
	defer cancel()

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go fn(cancelCtx, wg)
	wg.Wait()
	fmt.Println("Done...!")
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[fn] root-key :", ctx.Value("root-key"))

	//overriding the value set in the parent (hierarchy) context
	newRootValCtx := context.WithValue(ctx, "root-key", "new-root-value")

	// context created to share data from "fn"
	fnValCtx := context.WithValue(newRootValCtx, "fn-key", "fn-value")

	wg.Add(1)
	f1TimeoutCtx, f1Cancel := context.WithTimeout(fnValCtx, 5*time.Second)

	defer func() {
		fmt.Println("cancelling f1")
		f1Cancel()
	}()
	go f1(f1TimeoutCtx, wg)

	wg.Add(1)
	f2TimeoutCtx, f2Cancel := context.WithTimeout(fnValCtx, 6*time.Second)
	defer func() {
		fmt.Println("cancelling f2")
		f2Cancel()
	}()
	go f2(f2TimeoutCtx, wg)

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancellation signal received[fn]...")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("[fn] :", time.Now())
		}
	}
}

func f1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	// accessing the value set in the root (main)
	fmt.Println("[f1] root-key :", ctx.Value("root-key"))

	// accessing the value set in the fn
	fmt.Println("[f1] fn-key :", ctx.Value("fn-key"))

	i := 0
LOOP:
	for {
		i += 1
		select {
		case <-ctx.Done():
			fmt.Println("Cancellation signal received [f1]...")
			break LOOP
		default:
			time.Sleep(700 * time.Millisecond)
			fmt.Println("f1 :", i*2)
		}
	}
	fmt.Println("f1 completed")
}

func f2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	// accessing the value set in the root (main)
	fmt.Println("[f2] root-key :", ctx.Value("root-key"))

	// accessing the value set in the fn
	fmt.Println("[f2] fn-key :", ctx.Value("fn-key"))
	i := 0
LOOP:
	for {
		i += 1
		select {
		case <-ctx.Done():
			fmt.Println("Cancellation signal received [f2]...")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("f2 :", (i*2)+1)
		}
	}
	fmt.Println("f2 completed")
}
