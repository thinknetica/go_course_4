package main

import (
	"fmt"
)

// pow рекурсивно возводит "x" в степень "у"
func pow(x, y int) int {
	if y == 0 { // базовый случай
		return 1
	}
	return x * pow(x, y-1) // рекурсивный шаг
}

func main() {
	fmt.Println("3^2:", pow(3, 2)) // 3^2: 9
	fmt.Println("5^3:", pow(5, 3)) // 5^3: 125
}
