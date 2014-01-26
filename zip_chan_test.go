package pipe

import (
	"testing"
)

func TestZipChan(t *testing.T) {
	other := make(chan int, 5)
	in := make(chan int, 5)
	out := ZipChan(in, other).(chan []int)

	in <- 5
	in <- 10
	in <- 20
	other <- 6
	other <- 11
	close(other)

	count := 0
	for result := range out {
		count++

		expected := []int{count * 5, (count * 5) + 1}
		if len(result) != len(expected) {
			t.Fatal("expected channel output to match", expected, "but got", result)
		}

		for j := 0; j < len(result); j++ {
			if result[j] != expected[j] {
				t.Fatal("expected channel output to match", expected, "but got", result)
			}
		}
	}

	if count != 2 {
		t.Fatal("expected output to have 2 elements, but there were", count)
	}
}
