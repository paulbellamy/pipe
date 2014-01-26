package pipe

import (
	"testing"
)

func TestFlattenChan(t *testing.T) {
	in := make(chan []int, 5)
	out := FlattenChan(in).(chan int)

	in <- []int{1, 2}
	in <- []int{3, 4}
	in <- []int{5, 6}
	close(in)

	count := 0
	for result := range out {
		count++

		if result != count {
			t.Fatal("expected channel output to match", count, "but got", result)
		}
	}

	if count != 6 {
		t.Fatal("expected output to have 6 elements, but there were", count)
	}
}

func TestFlattenChanWhenAlreadyFlat(t *testing.T) {
	in := make(chan int, 5)
	out := FlattenChan(in).(chan int)

	in <- 1
	in <- 2
	in <- 3
	close(in)

	count := 0
	for result := range out {
		count++

		if result != count {
			t.Fatal("expected channel output to match", count, "but got", result)
		}
	}

	if count != 3 {
		t.Fatal("expected output to have 3 elements, but there were", count)
	}
}
