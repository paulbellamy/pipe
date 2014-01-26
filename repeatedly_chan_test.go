package pipe

import (
	"testing"
)

func TestRepeatedlyChan(t *testing.T) {
	count := -1
	out := RepeatedlyChan(func() int {
		count++
		return count
	}).(chan int)

	for i := 0; i < 5; i++ {
		result := <-out
		if result != i {
			t.Fatal("RepeatedlyChan was expected to output", i, "but output", result)
		}
	}
}
