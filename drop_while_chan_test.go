package pipe

import (
	"testing"
)

func TestDropWhileChan(t *testing.T) {
	in := make(chan int, 5)
	out := DropWhileChan(func(item int) bool {
		return item < 3
	}, in).(chan int)

	in <- 1
	in <- 2
	in <- 3
	in <- 2

	result := <-out
	if result != 3 {
		t.Fatal("DropWhileChan should have dropped all results until 3, but output", result)
	}

	result = <-out
	if result != 2 {
		t.Fatal("DropWhileChan should have kept all results after 3, but output", result)
	}

	close(in)
}
