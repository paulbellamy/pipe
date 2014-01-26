package pipe

import (
	"fmt"
)

func ExampleTakeWhileSlice() {
	in := []int{7, 4, 5, 6}
	out := TakeWhileSlice(func(item int) bool {
		return item != 5
	}, in).([]int)

	fmt.Println(out)

	// Output:
	// [7 4]
}
