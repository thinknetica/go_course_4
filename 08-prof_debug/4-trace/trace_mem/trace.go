package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
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

	nums := fillSlice()
	emptySlice(nums)

	runtime.GC()

	nums = fillSlice()
	emptySlice(nums)
}

func fillSlice() []int {
	var nums []int
	for i := 0; i < 100_000; i++ { // Изменить на 300_000.
		nums = append(nums, i)
	}
	return nums
}

func emptySlice(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		nums = append(nums[:i], nums[:i+1]...)
	}
}
