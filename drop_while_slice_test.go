package pipe

import (
	"testing"
)

func TestDropWhileSlice(t *testing.T) {
	in := []int{1,2,3,2}

  lessThan := func(x int) func(int) bool {
    return func(item int) bool {
      return item < x
    }
  }

	result := DropWhileSlice(lessThan(3), in).([]int)

  if len(result) != 2 {
		t.Fatal("DropWhileSlice should have dropped all results until 3, but output", result)
  }

  if result[0] != 3 {
		t.Fatal("DropWhile should have dropped all results until 3, but output", result)
  }

  if result[1] != 2 {
		t.Fatal("DropWhile should have dropped all results until 3, but output", result)
  }
}

func TestDropWhileSliceNeverPassing(t *testing.T) {
	in := []int{1,2,3,2}

  lessThan := func(x int) func(int) bool {
    return func(item int) bool {
      return item < x
    }
  }

	result := DropWhileSlice(lessThan(6), in).([]int)

  if len(result) != 0 {
		t.Fatal("DropWhileSlice should have dropped all results, but output", result)
  }
}
