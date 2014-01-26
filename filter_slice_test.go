package pipe

import (
	"fmt"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	even := func(item int) bool {
		return (item % 2) == 0
	}

	in := []int{7, 4}
	out := FilterSlice(even, in).([]int)

	if len(out) != 1 || out[0] != 4 {
		t.Fatal("FilterSlice(even, in) received 7 and 4 but output ", out)
	}
}

func TestFilterSliceTypeCoercion(t *testing.T) {
	long_enough := func(item fmt.Stringer) bool {
		return len(item.String()) > 1
	}

	in := []testStringer{7, 42}
	out := FilterSlice(long_enough, in).([]testStringer)

	if len(out) != 1 || out[0] != 42 {
		t.Fatal("FilterSlice(long_enough, in) received 7 and 42 but output ", out)
	}
}
