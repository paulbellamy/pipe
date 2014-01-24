package pipe

import (
	"fmt"
	"testing"
)

func sum(sum, item int) int {
	return sum + item
}

func TestReduceSlice(t *testing.T) {
	in := []int{5, 10, 20}
	out := ReduceSlice(sum, 0, in).(int)

	if out != 35 {
		t.Fatal("ReduceSlice(sum, 0, []int{5, 10, 20}) output ", out)
	}
}

func appendToString(str string, item fmt.Stringer) string {
  return fmt.Sprintf("%s%s", str, item.String())
}

func TestReduceSliceTypeCoercion(t *testing.T) {
	in := []testStringer{1,2,3}
	out := ReduceSlice(appendToString, "a", in).(string)

	if out != "a123" {
		t.Fatal("ReduceSlice(appendToString, \"a\", []int{1, 2, 3}) output ", out)
	}
}
