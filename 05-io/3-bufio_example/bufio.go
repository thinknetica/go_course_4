package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./bufio.go")
	if err != nil {
		log.Fatal(err)
	}
	b, err := get(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	f, err = os.Create("./bufio_copy.go")
	if err != nil {
		log.Fatal(err)
	}
	err = store(f, b)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}

func store(w io.Writer, b []byte) error {
	// Здесь должна быть некотороая логика.
	// После обработки данных мы их записываем туда,
	// куда хочет вызывающий код.
	_, err := w.Write(b)
	return err
}

func get(r io.Reader) ([]byte, error) {
	// буфер уже является частью сканера
	scanner := bufio.NewScanner(r)
	var b []byte
	for scanner.Scan() {
		b = append(b, []byte(scanner.Text()+"\n")...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// Здесь какая-то логика.
	return b, nil
}
