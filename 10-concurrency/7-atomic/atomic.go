package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var c atomic.Int64

	N := 100_000

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			c.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Счётчик:", c.Load())
}
