package pipe

import (
	"testing"
)

func TestForEachChan(t *testing.T) {
	count := 0

	in := make(chan int, 5)
	counter := func(item int) {
		count++
	}
	out := ForEachChan(counter, in).(chan int)

	in <- 5
	in <- 6
	in <- 7

	// drain the pipe
	for i := 5; i <= 7; i++ {
		result := <-out
		if result != i {
			t.Fatal("counting ForEachChan modified ", i, " into ", result)
		}
	}

	if count != 3 {
		t.Fatal("counting ForEachChan received 3 items but counted ", count, "/ 3")
	}

	close(in)
}
