package main

import (
	"fmt"

	"github.com/DmitriyVTitov/size"
)

type example struct {
	a []int
	b bool
	c int32
	d string
}

func main() {
	ex := example{
		a: []int{1, 2, 3}, // ?
		b: true,           // ?
		d: "1234",         // ?
	} // ?
	fmt.Println("Размер в байтах для ex:", size.Of(ex))
	// Как получается результат?

	ex1 := example{
		a: []int{1, 2, 3}, // ?
		b: true,           // ?
		d: "1234",         // ?
		c: 100,
	} // ?
	fmt.Println("Размер в байтах для ex1:", size.Of(ex1))
	// Как получается результат?
}
