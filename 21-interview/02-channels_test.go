package interview

import (
	"testing"
)

func TestSelect(t *testing.T) {
	Select()
}

func TestProduceConsume(t *testing.T) {
	ProduceConsume()
}

func TestFanIn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()
	go func() {
		for i := 11; i < 20; i++ {
			ch1 <- i
		}
	}()

	ch := FanIn(ch1, ch2)
	for val := range ch {
		t.Log(val)
	}
}
