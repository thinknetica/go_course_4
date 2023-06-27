package names

import "testing"

const testStr = "ABCdE"

func Test_index(t *testing.T) {
	got := index(testStr, 'd')
	want := 3
	if got != want {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
