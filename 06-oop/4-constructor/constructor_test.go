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

func TestGuitar(t *testing.T) {
	var g Guitar
	g.Tabs = make(map[string]string)
	g.Tabs["Metallica"] = "One"
}
