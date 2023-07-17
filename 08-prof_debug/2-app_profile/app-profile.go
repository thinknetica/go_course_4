package main

import (
	"math"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go compute1(1_000_000, 1_000_000)
	go compute2(2_000_000, 2_000_000)

	go mem1()
	go mem2()

	http.ListenAndServe(":80", nil)
}

func compute1(n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_ = math.Sin(float64(i)) * math.Sin(float64(j))
		}
	}
}

func compute2(n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_ = math.Cos(float64(i)) * math.Cos(float64(j))
		}
	}
}

func mem1() {
	var arr []int32
	for i := 0; i < 10_000_000; i++ {
		arr = append(arr, int32(i))
	}
	_ = arr
	select {}
}

func mem2() {
	var arr []int64
	for i := 20_000_000; i > 0; i-- {
		arr = append(arr, int64(i))
	}
	_ = arr
	select {}
}
