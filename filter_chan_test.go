package pipe

import (
	"fmt"
	"testing"
)

func TestFilterChan(t *testing.T) {
	even := func(item int) bool {
		return (item % 2) == 0
	}

	in := make(chan int)
	out := FilterChan(even, in).(chan int)

	go func() {
		in <- 7
		in <- 4
		in <- 5
		in <- 2
		close(in)
	}()

	if result := <-out; result != 4 {
		t.Fatal("FilterChan(even, in) received 7,4,5,2, but output ", result)
	}

	if result := <-out; result != 2 {
		t.Fatal("FilterChan(even, in) received 7,4,5,2, but output ", result)
	}

	if _, ok := <-out; ok {
		t.Fatal("FilterChan(even, in) wasn't closed after in was closed")
	}
}

func TestFilterChanTypeCoercion(t *testing.T) {
	long_enough := func(item fmt.Stringer) bool {
		return len(item.String()) > 1
	}

	in := make(chan testStringer)
	out := FilterChan(long_enough, in).(chan testStringer)

	go func() {
		in <- 7
		in <- 42
	}()

	if result := <-out; result != 42 {
		t.Fatal("FilterChan(long_enough, in) received 7 and 42 but output ", out)
	}

	close(in)
}
