package main

import (
	"fmt"

	"github.com/DmitriyVTitov/size"
)

const N = 1_000_000_000

var slice = make([]int, N)

func main() {
	newSlice := make([]int, N+1)
	copy(newSlice, slice)
	newSlice[len(newSlice)-1] = 25

	s := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		s[i] = i
	}
	_ = s

	newSlice[1] = 10
	x, y := size.Of(slice), size.Of(newSlice)
	fmt.Printf("Size of 'slice' is %vGB\n", x/1024/1024/1024)
	fmt.Printf("Size of 'newSlice' is %vGB\n", y/1024/1024/1024)
	fmt.Printf("Size of both slices is %vGB\n", (x+y)/1024/1024/1024)

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:3] // {2, 3}
	fmt.Println("s2:", s2)
	fmt.Println("s1 before:", s1)
	s2[0] = 10
	fmt.Println("s1 after:", s1)
}
