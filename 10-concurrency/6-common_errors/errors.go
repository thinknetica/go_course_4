package main

import (
	"fmt"
	"time"
)

func error1() {
	// Ошибка №1. Доступ к итератору цикла через замыкание.
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1)
}

func error2() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	// Ошибка №2. Использование таймера для синхронизации.
	time.Sleep(1)
}

func error3() {
	f := func(i int) {
		fmt.Println(i)
	}
	// Ошибка №3. Отсутствие контроля за завершением потоков.
	for i := 0; i < 10; i++ {
		go f(i)
	}
}

func main() {
	//error1()
	//error2()
	//error3()
}
