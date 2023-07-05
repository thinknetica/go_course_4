package main

import (
	"io"
	"log"
	"os"
)

// sendReversedString разворачивает строку и отправляет её получателю.
func sendReversedString(s string, w io.Writer) error {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	_, err := w.Write([]byte(string(runes)))
	return err
}

func main() {
	f, err := os.Create("./strings.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Вывод на примере файла.
	err = sendReversedString("АБЫРВАЛГ", f)
	if err != nil {
		log.Fatal(err)
	}

	// Вывод на примере консоли.
	err = sendReversedString("АБЫРВАЛГ", os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func get(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
	/*/
	var b []byte
	var buf = make([]byte, 10)
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return nil, err
		}
		b = append(b, buf...)
	}
	// Здесь какая-то логика.
	return b, nil*/
}
