package pipe

import (
	"fmt"
	"testing"
)

func TestTakeWhileSlice(t *testing.T) {
	in := []int{7, 4, 5, 6}
	out := TakeWhileSlice(func(item int) bool {
		return item != 5
	}, in).([]int)

	if len(out) != 2 || out[0] != 7 || out[1] != 4 {
		t.Fatal("takewhile should have returned [7 4], but returned", out)
	}
}

func TestTakeWhileSliceTypeCoercion(t *testing.T) {
	strLenLessThan := func(length int) func(fmt.Stringer) bool {
		return func(x fmt.Stringer) bool {
			return len(x.String()) < length
		}
	}

	in := []testStringer{8, 9, 10, 9}
	out := TakeWhileSlice(strLenLessThan(2), in).([]testStringer)

	if len(out) != 2 || out[0] != 8 || out[1] != 9 {
		t.Fatal("Expected:", []testStringer{8, 9}, "\nGot:", out)
	}
}
