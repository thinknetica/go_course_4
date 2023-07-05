package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(f.Name(), []byte("Текст"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Данные файла:\n%s\n", data)

	// вариант чтения с помощью io.Reader
	var file *os.File
	file, err = os.Open(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// что здесь не правильно?
	buf := make([]byte, 6)
	n, err := file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	buf = buf[:n]
	fmt.Printf("Данные файла:\n%s\n", buf)
}
