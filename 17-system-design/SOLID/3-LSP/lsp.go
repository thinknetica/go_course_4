package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Пакет демонстрирует выполнение принципа LSP в Go.

// Интерфейс с контрактом на представление
// значения некоторого типа данных в виде
// последовательности байт.
type Serializer interface {
	Serialize() []byte
}

type String string

func (s String) Serialize() []byte {
	return []byte(s)
}

type Bool bool

func (b Bool) Serialize() []byte {
	return []byte(fmt.Sprintf("%v", b))
}

// Музыкальный альбом.
type Album struct {
	Title String
	Year  uint
}

// Реализация контракта.
func (a Album) Serialize() []byte {
	b, err := json.Marshal(a)
	if err != nil {
		return nil
	}
	return b
}

// Функция принимает на вход объект, в который требуется
// сохранить состояние: файл, принтер и т.д.
// Также принимается любое количество объектов, поддерживающих
// сериализацию.
func WriteObject(w io.Writer, objects ...Serializer) error {
	for _, obj := range objects {
		b := obj.Serialize()
		_, err := w.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var s String = "ABC"
	var b Bool = true
	var a = Album{Title: "Nevermind", Year: 1991}

	f, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = WriteObject(f, s, b, a)
	if err != nil {
		log.Fatal(err)
	}
}
