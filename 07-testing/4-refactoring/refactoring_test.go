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
