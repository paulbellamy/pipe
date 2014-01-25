package pipe

import (
  "testing"
)

func TestDropSlice(t *testing.T) {
	in := []int{0, 1, 2, 3, 4}
	out := DropSlice(3, in).([]int)

	if len(out) != 2 || out[0] != 3 || out[1] != 4 {
		t.Fatal("DropSlice(3) expected", []int{3, 4}, "but output ", out)
	}
}

func TestDropSliceNone(t *testing.T) {
	in := []int{0, 1, 2, 3, 4}
	out := DropSlice(0, in).([]int)

	if len(out) != 5 {
		t.Fatal("DropSlice(3) expected", in, "but output ", out)
	}
}

func TestDropSliceEmpty(t *testing.T) {
	in := []int{}
	out := DropSlice(3, in).([]int)

	if len(out) != 0 {
		t.Fatal("DropSlice(3) expected", in, "but output ", out)
	}
}

func TestDropSliceMoreThanRemains(t *testing.T) {
	in := []int{0, 1, 2, 3, 4}
	out := DropSlice(6, in).([]int)

	if len(out) != 0 {
		t.Fatal("DropSlice(3) expected", []int{}, "but output ", out)
	}
}
