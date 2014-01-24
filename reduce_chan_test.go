package pipe

import (
	"testing"
)

func TestReduceChan(t *testing.T) {
	in := make(chan int, 5)

	go func() {
		in <- 5
		in <- 10
		in <- 20
		close(in)
	}()

	out := ReduceChan(sum, 0, in).(int)

	if out != 35 {
		t.Fatal("ReduceChan(sum, 0, []int{5, 10, 20}) output ", out)
	}
}
