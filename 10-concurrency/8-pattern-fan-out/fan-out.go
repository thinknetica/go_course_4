package main

import (
	"fmt"
	"sync"
)

func process(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * val
	}
}

func main() {
	src := make(chan int)
	res := make(chan int)

	// Запуск рабочих потоков.
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			process(src, res)
		}()
	}

	// Поток с обработкой результатов.
	go func() {
		for val := range res {
			fmt.Println(val)
		}
	}()

	// Поток с заданиями.
	go func() {
		for i := 0; i < 10; i++ {
			src <- i
		}
		close(src)
	}()

	wg.Wait()
	close(res)
}
