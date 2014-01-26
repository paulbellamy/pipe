package pipe

import (
	"fmt"
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

func TestTakeWhileChanTypeCoercion(t *testing.T) {
	strLenLessThan := func(length int) func(fmt.Stringer) bool {
		return func(x fmt.Stringer) bool {
			return len(x.String()) < length
		}
	}

	in := make(chan testStringer, 5)
	out := TakeWhileChan(strLenLessThan(2), in).(chan testStringer)

	in <- 8
	in <- 9
	in <- 10
	in <- 9

	if result := <-out; result != 8 {
		t.Fatal("Expected:", 8, "\nGot:", result)
	}

	if result := <-out; result != 9 {
		t.Fatal("Expected:", 9, "\nGot:", result)
	}

	if _, ok := <-out; ok {
		t.Fatal("takewhile pipe should have closed the channel after turning it off")
	}

	close(in)
}
