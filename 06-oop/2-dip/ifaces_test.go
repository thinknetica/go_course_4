package main

import (
	"testing"
)

func Test_logMsg(t *testing.T) {
	s := New(&MemLogger{})
	_ = s

	l := new(CustomLogger)
	err := logMsg(l, "msg")
	if err != nil {
		t.Fatal(err)
	}
}
