package refactoring

import (
	"testing"
)

func Test_rublesToUSD(t *testing.T) {
	const rubles = 1000
	res, err := rublesToUSD(rubles)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v рублей - это %.2f долларов\n", rubles, res)
}

func Test_calc(t *testing.T) {
	want := 10
	got := calc(1000, 100)
	if got != float64(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
