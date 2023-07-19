package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"
)

// serializer - абстрактный интерфейсный тип даных.
// Объявляет контракт на сериализацию.
type serializer interface {
	serialize() ([]byte, error) // контракт интерфейса
}

// Конкретный тип данных.
type person struct {
	Name string
	Age  int
}

// Метод, выполняющий контракт интерфейса.
func (p *person) serialize() ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//
// Другой тип, выполняющий контракт интерфейса.
//

type car struct {
	Modle string
	Year  int
}

func (c *car) serialize() ([]byte, error) {
	b, err := xml.Marshal(c)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Полиморфическая функция, принимающая интерфейс.
func store(s serializer, w io.Writer) error {
	b, err := s.serialize()
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func main() {
	p := person{
		Name: "Курт",
		Age:  27,
	}
	var s serializer = &p

	f, err := os.Create("./output.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = store(s, f)
	if err != nil {
		log.Fatal(err)
	}

	err = store(&p, f)
	if err != nil {
		log.Fatal(err)
	}

	// *************************************************** //

	c := car{
		Modle: "Beetle",
		Year:  1970,
	}

	err = store(&c, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
