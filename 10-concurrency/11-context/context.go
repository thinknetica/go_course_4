package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, in <-chan string, out chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("объект контекста вызвал завершение потока!")
			close(out)
			fmt.Println("Val:", ctx.Value("key"))
			return
		case val := <-in:
			fmt.Println("Обработка сообщения:", val)
			out <- fmt.Sprintf("Сообщение %s обработано", val)
		}
	}
}

func main() {
	data := make(chan string)
	result := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 700*time.Millisecond)
	ctx = context.WithValue(ctx, "key", "val")
	defer cancel()

	go worker(ctx, data, result)

	go func() {
		for i := 0; i < 10; i++ {
			if i > 8 {
				cancel()
				break
			}
			data <- fmt.Sprintf("сообщение #%d", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for val := range result {
		fmt.Println(val)
	}
}
