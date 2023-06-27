package main

import (
	"testing"
	"time"
)

func Test_unbufChan(t *testing.T) {
	ch := unbufChan()
	for val := range ch {
		t.Log(val)
		time.Sleep(time.Second)
	}
}

func Test_bufChan(t *testing.T) {
	ch := bufChan()
	for val := range ch {
		t.Log(val)
		time.Sleep(time.Second)
	}
}
