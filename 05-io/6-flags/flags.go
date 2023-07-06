package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("filename", "file.txt", "имя файла")

	var n int
	flag.IntVar(&n, "n", 10, "количество")
	flag.IntVar(&n, "number", 10, "количество")

	flag.Parse()

	flag.PrintDefaults()

	fmt.Println(*s, n)
}
