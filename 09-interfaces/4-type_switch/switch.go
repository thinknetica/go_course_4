package main

import (
	"fmt"
)

func main() {
	arr := []interface{}{10, "String", true}
	for _, v := range arr {
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
