package interview

import (
	"fmt"
	"sync"
)

func Select() {
	ch := make(chan string)
	select {
	case ch <- "string":
		fmt.Println("получилось записать в канал")
	default:
		fmt.Println("не получилось записать в канал")
	}
}

func ProduceConsume() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for val := range ch {
		fmt.Println(val)
	}
}

func FanIn[T any](channels ...<-chan T) <-chan T {
	ch := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func(in <-chan T) {
			defer wg.Done()
			for i := range in {
				ch <- i
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
