package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	/*for {
		val, ok := <-ch
		fmt.Println(val, ok)
		if !ok {
			break
		}
	}*/

	for val := range ch {
		fmt.Println(val)
	}
}
