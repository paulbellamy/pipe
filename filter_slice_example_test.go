package pipe

import (
	"fmt"
)

func ExampleFilterSlice() {
	even := func(item int) bool {
		return (item % 2) == 0
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := FilterSlice(even, numbers).([]int)

  fmt.Println(out)

  // Output:
  // [2 4 6 8 10]
}
