package main

import (
	"fmt"
)

func unbufChan() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		fmt.Println("Функция завершила работу!")
	}()
	return ch
}

func bufChan() chan int {
	ch := make(chan int, 10)
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		fmt.Println("Функция завершила работу!")
	}()
	return ch
}
