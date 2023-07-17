package bsearch

// Binary возвращает номер элемента в массиве или -1. Используется бинарный поиск.
func Binary(data []int, item int) int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == item {
			return mid
		}
		if data[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Simple возвращает номер элемента в массиве или -1. Используется простой поиск.
func Simple(data []int, item int) int {
	for i := range data {
		if data[i] == item {
			return i
		}
	}
	return -1
}
