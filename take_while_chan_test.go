package pipe

import (
	"testing"
)

func TestTakeWhileChan(t *testing.T) {
	take := true
	in := make(chan int, 5)
	out := TakeWhileChan(func(item int) bool {
		return take
	}, in).(chan int)

	in <- 7
	in <- 4
	take = false
	in <- 5
	in <- 6

	<-out
	<-out
	if _, ok := <-out; ok {
		t.Fatal("takewhile pipe should have closed the channel after turning it off")
	}

	close(in)
}

func TestTakeWhileChanTypeCoerciton(t *testing.T) {
	t.Fatal("Not Implemented")
}
