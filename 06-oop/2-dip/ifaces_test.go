package main

import (
	"testing"
)

func Test_logMsg(t *testing.T) {
	l := new(CustomLogger)
	err := logMsg(l, "msg")
	if err != nil {
		t.Fatal(err)
	}
}
