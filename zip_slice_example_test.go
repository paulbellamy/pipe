package pipe

import (
	"fmt"
)

func ExampleZipSlice() {
	in := []int{5, 10, 20}
	other := []int{6, 11}
	out := ZipSlice(in, other).([][]int)

	for _, result := range out {
		fmt.Println(result)
	}

	// Output:
	// [5 6]
	// [10 11]
}
