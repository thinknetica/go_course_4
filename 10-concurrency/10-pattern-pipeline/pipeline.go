package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func proc(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func proc2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n * n
		}
		close(out)
	}()
	return out
}

func main() {
	src := gen(1, 2, 3, 4, 5)
	res := proc(proc2(proc(proc(src))))

	for val := range res {
		fmt.Println(val)
	}
}
