package main

import (
	"fmt"
)

func main() {
	slice := []any{10, "String", true}
	for _, v := range slice {
		switch v.(type) {
		case int:
			fmt.Println("Число:", v)
		case string:
			fmt.Println("Строка:", v)
		default:
			fmt.Println("Неизвестно:", v)
		}
	}
}
