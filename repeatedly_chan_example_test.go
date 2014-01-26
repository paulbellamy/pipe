package pipe

import (
	"fmt"
)

func ExampleRepeatedlyChan() {
	count := -1
	out := RepeatedlyChan(func() int {
		count++
		return count
	}).(chan int)

	for i := 0; i < 3; i++ {
		fmt.Println(<-out)
	}

	// Output:
	// 0
	// 1
	// 2
}
