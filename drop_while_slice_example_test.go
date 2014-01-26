package pipe

import (
  "fmt"
)

func ExampleDropWhileSlice() {
  in := []int{1,2,3,2}

  lessThan := func(x int) (func (int) bool) {
    return func(item int) bool {
      return item < 3
    }
  }

	out := DropWhileSlice(lessThan(3), in).([]int)

  fmt.Println(out)

  // Output:
  // [3 2]
}
