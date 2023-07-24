package main

import (
	"fmt"
	"log"
)

type modern interface {
	isModern() bool
}

type old interface {
	isOld() bool
	isObsolete() bool
}

type cassete struct {
	label string
}

func (c cassete) isModern() bool {
	return false
}

func (c cassete) isOld() bool {
	return true
}

func (c *cassete) isObsolete() bool {
	return false
}

func main() {
	var m modern
	var c cassete

	m = c

	//m.label = "Ace Of Base" - ошибка
	m.isModern()

	switch c1 := m.(type) {
	case cassete:
		c1.label = "Ace Of Base"
	case old:
		c1.isOld()
	}

	if casseteInstance, ok := m.(cassete); ok {
		casseteInstance.label = "Ace Of Base"
		log.Printf("%+v", casseteInstance)
	}

	if oldInstance, ok := m.(old); ok {
		fmt.Println("Old: ", oldInstance.isOld(), "\tObsolete: ", oldInstance.isObsolete())
	}
}
