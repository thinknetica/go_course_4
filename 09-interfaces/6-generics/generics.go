package main

import "fmt"

type ordered interface {
	int | float64 | float32 | byte | string
}

// min - обобщённая функция, возврающая минимальный элемент в массиве.
func min[T ordered](in []T) (T, bool) {
	var res T
	if len(in) == 0 {
		return res, false
	}
	res = in[0]
	for _, v := range in {
		if v < res {
			res = v
		}
	}
	return res, true
}

func main() {
	ints := []int{-2, 0, 2, 3}
	floats := []float32{2.3, 3.4, 4.5}
	strings := []string{"def", "Абв", "ABC"}

	if val, ok := min[int](ints); ok {
		fmt.Printf("%v\n", val)
	}

	if val, ok := min(floats); ok {
		fmt.Printf("%v\n", val)
	}

	if val, ok := min(strings); ok {
		fmt.Printf("%v\n", val)
	}
}
