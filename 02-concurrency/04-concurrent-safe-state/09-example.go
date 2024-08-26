package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Add(delta int) {
	c.Lock()
	{
		c.count += delta
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count : ", counter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Add(1)
}
