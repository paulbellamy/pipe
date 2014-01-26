package pipe

import (
	"fmt"
)

func ExampleFilterChan() {
	even := func(item int) bool {
		return (item % 2) == 0
	}

	numbers := make(chan int)
	out := FilterChan(even, numbers).(chan int)

	go func() {
		numbers <- 1
		numbers <- 2
		numbers <- 3
		numbers <- 4
		numbers <- 5
		numbers <- 6
		numbers <- 7
		numbers <- 8
		numbers <- 9
		numbers <- 10
		close(numbers)
	}()

	// Print each output
	for result := range out {
		fmt.Println(result)
	}

	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}
