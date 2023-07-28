package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generator(ch chan<- string, num int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		ch <- "Сообщение из канала №" + strconv.Itoa(num)
	}
	close(ch)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan string)
	ch2 := make(chan string)
	go generator(ch1, 1)
	go generator(ch2, 2)

OUT:
	for {
		select {
		case val, ok := <-ch1:
			if !ok {
				break OUT
			}
			fmt.Println(val)
		case val := <-ch2:
			fmt.Println(val)
		default:
			continue
		}
	}
}
