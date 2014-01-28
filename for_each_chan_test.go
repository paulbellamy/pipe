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

	in <- 5
	in <- 6
	in <- 7
	close(in)

	ForEachChan(counter, in)

	if count != 3 {
		t.Fatal("counting ForEachChan received 3 items but counted ", count, "/ 3")
	}
}
