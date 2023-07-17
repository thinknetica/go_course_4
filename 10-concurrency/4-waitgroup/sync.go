package main

import (
	"fmt"
	"sync"
)

func printN(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n)
}

func step1(ch chan<- string) {
	ch <- "message from Step 1"
}
func step2(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	const N = 10
	var wg sync.WaitGroup
	wg.Add(N)
	// Запуск отдельных потоков.
	for i := 0; i < N; i++ {
		go printN(i, &wg)
	}
	// Ожидание завершения потоков.
	wg.Wait()

	// Синхронизация двух потоков с помощью
	// сообщений в канале.
	var ch = make(chan string)
	go step1(ch)
	step2(ch)
}
