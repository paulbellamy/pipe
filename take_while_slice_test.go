package pipe

import (
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
