package flags

import (
	"flag"
	"testing"
)

var f = flag.Bool("int", false, "")

func TestFunc(t *testing.T) {
	flag.Parse()
	if *f {
		t.Log("FLAG")
	}
	if !*f {
		t.Log("NO FLAG")
	}
}
