package pipe

import (
	"fmt"
	"testing"
)

func TestMapSlice(t *testing.T) {
	count := 0
	counter := func(item int) string {
		count++
		return fmt.Sprint(count)
	}
	in := []int{7, 4, 5}
	out := MapSlice(counter, in).([]string)

	for i := 1; i <= 3; i++ {
		if result := out[i-1]; result != fmt.Sprint(i) {
			t.Fatal("MapSlice received", in, "but output", out, "expected", []string{"1", "2", "3"})
		}
	}
}

func TestMapSliceTypeCoercion(t *testing.T) {
	count := 0
	counter := func(item fmt.Stringer) string {
		count++
		return fmt.Sprint(count)
	}
	in := []testStringer{7, 4, 5}
	out := MapSlice(counter, in).([]string)

	for i := 1; i <= 3; i++ {
		if result := out[i-1]; result != fmt.Sprint(i) {
			t.Fatal("MapSlice received", in, "but output", out, "expected", []string{"1", "2", "3"})
		}
	}
}
