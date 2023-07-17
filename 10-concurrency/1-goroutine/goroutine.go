package main

import (
	"fmt"
	"time"
)

func printN(n int) int {
	fmt.Println(n)
	return n * 2
}

func main() {
	for i := 0; i < 10; i++ {
		go printN(i) // что будет выведено на экран?
	}

	time.Sleep(time.Second)
}
