package main

import (
	"fmt"
)

// Course - курс по Go.
type Course struct {
	author string
}

// SetAuthor устанавливает имя автора.
func (c *Course) SetAuthor(name string) {
	c.author = name
}

// Author возвращает имя атора курса.
func (c Course) Author() string {
	return c.author
}

func main() {
	c := new(Course) // new создаёт переменную и возвращает указатель на неё
	c.SetAuthor("1")
	fmt.Println(c.Author())

	var course Course
	course.SetAuthor("2")
	fmt.Println(course.Author()) // Что будет выведено здесь?

	// The rule about pointers vs. values for receivers is that value methods
	// can be invoked on pointers and values,
	// but pointer methods can only be invoked on pointers.
}

// *** Вопрос: что будет выведено, если получатели методов сделать значениями?
