package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("./file.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	err = os.WriteFile(f.Name(), []byte("Текст"), 0666)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := os.ReadFile(f.Name())
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Данные файла:\n%s\n", data)

	// вариант чтения с помощью io.Reader
	var file *os.File
	file, err = os.Open(f.Name())
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// что здесь не правильно?
	buf := make([]byte, 6)
	n, err := file.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	buf = buf[:n]
	fmt.Printf("Данные файла:\n%s\n", buf)
}
