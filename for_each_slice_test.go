package pipe

import (
	"testing"
)

func TestForEachSlice(t *testing.T) {
	count := 0

	in := []int{5, 6, 7}
	counter := func(item int) {
		count++
	}
	ForEachSlice(counter, in)

	if count != 3 {
		t.Fatal("counting ForEachSlice received 3 items but counted ", count, "/ 3")
	}
}
