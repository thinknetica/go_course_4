package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type T struct {
	I int
	S string
	B bool
}

func main() {

	// сериализация
	t1 := T{I: 10, S: "ABC", B: true}
	b, err := json.Marshal(t1) // установить точку останова
	if err != nil {
		log.Fatal(err)
	}

	// десериализация
	var t2 T
	err = json.Unmarshal(b, &t2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", t2)
	fmt.Println(reflect.DeepEqual(t1, t2))
}
