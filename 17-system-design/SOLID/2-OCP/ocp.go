package ocp

import "fmt"

// Пакет демонстрирует типовое нарушение
// принципа OCP.
// Функции занимаются анализом и обработкой
// результата вместо того, чтобы вернуть
// его вызывающему коду.
// В результате у пользователя пакета появляется
// потребность изменить код пакета, что является
// нарушением принципа OCP.

func Avg(nums []int) {
	if len(nums) == 0 {
		return
	}

	var num float64
	var sum int
	for _, n := range nums {
		sum += n
	}

	num = float64(sum / len(nums))

	// следует вернуть результат, а не печатать его
	fmt.Println(num)
}

func Max(nums []int) int {
	// функция проанализировала входные данные и
	// вернула нелогичный результат, хотя для этого
	// нет никаких оснований
	if len(nums) == 0 {
		return -1
	}

	var num int
	for _, n := range nums {
		if n > num {
			num = n
		}
	}

	return num
}
