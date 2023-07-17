package main

import (
	"log"
	"math"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	// файл, в который будет сохраняться трассировка
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	// запуск трассировки и отложенное завершение
	trace.Start(f)
	defer f.Close()
	defer trace.Stop()

	go compute1()
	go compute1()
	go compute2()
	go compute2()

	time.Sleep(time.Second * 1)
}

func compute1() {
	for {
		_ = math.Sin(1) * math.Cos(10)
	}
}

func compute2() {
	for {
		_ = math.Tan(1)
	}
}
