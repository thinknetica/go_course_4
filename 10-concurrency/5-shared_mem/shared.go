package main

import (
	"fmt"
	"sync"
)

// add прибавляет x к sum
func add(sum *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 1; i <= 10_000; i++ {
		mu.Lock()
		*sum++
		mu.Unlock()
	}
}

func main() {
	sum := 0
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go add(&sum, &wg, &mu)
	go add(&sum, &wg, &mu)

	wg.Wait()

	fmt.Println(sum) // Что будет на экране?
}
