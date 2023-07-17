package main

import (
	"encoding/json"
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

// Полиморфическая функция, принимающая интерфейс.
func str(s serializer) string {
	b, err := s.serialize()
	if err != nil {
		return ""
	}
	return string(b)
}

func main() {
	p := &person{
		Name: "Курт",
		Age:  27,
	}
	var s serializer
	s = p
	print(str(s))
}
