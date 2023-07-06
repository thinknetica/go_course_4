package guitar

import (
	"testing"
)

func TestNew(t *testing.T) {
	g := New("Fender", "Stratocaster")
	got := g.Info()
	want := "Гитара Fender Stratocaster"
	if got != want {
		t.Errorf("получили %s, ожидалось %s", got, want)
	}
}
