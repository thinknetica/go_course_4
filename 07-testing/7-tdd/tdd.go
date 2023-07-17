package tdd

func Fact(n int) int {
	if n < 1 {
		return 1
	}

	return n * Fact(n-1)
}
