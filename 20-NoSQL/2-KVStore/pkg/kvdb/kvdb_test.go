package kvdb

import "testing"

func TestDB_GET(t *testing.T) {
	db := New()
	want := "b"
	db.SET("a", want)
	got := db.GET("a")
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	got = db.GET("!@#")
	if got != "" {
		t.Fatalf("got %s, want %s", got, "")
	}
}
