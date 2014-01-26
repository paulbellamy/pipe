package pipe

import (
	"fmt"
)

func ExampleMapCatSlice() {
	printer := func(item int) []string {
		return []string{fmt.Sprintf("%d", item), fmt.Sprintf("%.3d", item)}
	}
	in := []int{7, 4, 5}
	result := MapCatSlice(printer, in).([]string)

	for _, item := range result {
		fmt.Println(item)
	}

	// Output:
	// 7
	// 007
	// 4
	// 004
	// 5
	// 005
}
