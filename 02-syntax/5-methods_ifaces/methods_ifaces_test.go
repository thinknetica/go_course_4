package main

import "testing"

func Test_myStr_print(t *testing.T) {
	var s myStr = "ABC"
	var p printer
	p = &s
	p.print()
	t.Log(s)
}
