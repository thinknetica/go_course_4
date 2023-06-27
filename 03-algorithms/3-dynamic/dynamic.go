package main

import (
	"fmt"
	"log"
	"strconv"
)

// fibo рекурсивно вычисляет числа фибоначчи, используя идею динамического программирования
func fibo(n int, cache map[int]int) int {
	// проверка наличия уже вычисленного результата в кэше
	if val, ok := cache[n]; ok {
		return val
	}
	// базовая часть рекурсии
	if n < 2 {
		return 1
	}
	// рекурсивный шаг с записью в кэш
	num := fibo(n-1, cache) + fibo(n-2, cache)
	cache[n] = num
	return num
}

func main() {
	cache := make(map[int]int)

	for {
		fmt.Print("Введите число: ")
		var in string
		_, err := fmt.Scanln(&in)
		if err != nil {
			log.Fatal(err)
		}
		n, err := strconv.Atoi(in)
		if err != nil {
			log.Fatal(err)
		}
		num := fibo(n, cache)
		fmt.Printf("Число Фибоначчи №%d = %d\n", n, num)
	}
}
