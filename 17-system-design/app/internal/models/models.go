package models

type Book struct {
	Name   string
	Author string
}

var Books = []Book{
	{
		Name:   "The Lord Of The Rings",
		Author: "J.R.R. Tolkien",
	},
}
