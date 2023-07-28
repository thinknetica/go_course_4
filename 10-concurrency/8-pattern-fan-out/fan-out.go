package main

import (
	"fmt"
	"runtime"
	"sync"
)

func processor(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * val
	}
}

func main() {
	src := make(chan int)
	res := make(chan int)

	n := runtime.NumCPU()
	// Запуск рабочих потоков.
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			processor(src, res)
		}()
	}

	// Поток с заданиями.
	go func() {
		for i := 0; i < 10; i++ {
			src <- i
		}
		close(src)
	}()

	// Поток с обработкой результатов.
	go func() {
		for val := range res {
			fmt.Println(val)
		}
	}()

	wg.Wait()

	close(res)
}
