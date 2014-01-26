package pipe

import (
	"testing"
)

func TestReverseChan(t *testing.T) {
	in := make(chan int, 5)
	out := ReverseChan(in).(chan int)

	in <- 7
	in <- 6
	in <- 5
	close(in)

	// drain the pipe
	for i := 5; i <= 7; i++ {
		result := <-out
		if result != i {
			t.Fatal("ReverseChan modified ", i, " into ", result)
		}
	}

	if _, ok := <-out; ok {
		t.Fatal("ReverseChan received 3 items but returned more")
	}
}
