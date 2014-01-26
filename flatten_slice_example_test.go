package pipe

import (
	"fmt"
)

func ExampleFlattenSlice() {
	in := [][]int{{1, 2}, {3, 4}, {5}}
	out := FlattenSlice(in).([]int)

	for _, result := range out {
		fmt.Println(result)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
