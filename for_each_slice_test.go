package pipe

import (
	"testing"
)

func TestForEachSlice(t *testing.T) {
	count := 0

	in := []int{5,6,7}
	counter := func(item int) {
		count++
	}
	out := ForEachSlice(counter, in).([]int)

	// drain the pipe
  for i := 0; i < len(in); i++ {
		if out[i] != in[i] {
			t.Fatal("counting ForEachSlice modified ", in[i], " into ", out[i])
		}
  }

	if count != 3 {
		t.Fatal("counting ForEachSlice received 3 items but counted ", count, "/ 3")
	}
}
