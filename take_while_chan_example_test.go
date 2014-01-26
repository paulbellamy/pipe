package pipe

import (
	"fmt"
)

func ExampleTakeWhileChan() {
	in := make(chan int, 5)
	out := TakeWhileChan(func(item int) bool {
		return item != 5
	}, in).(chan int)

	in <- 7
	in <- 4
	in <- 5
	in <- 6

	for result := range out {
		fmt.Println(result)
	}

	fmt.Println("Closed!")

	// Output:
	// 7
	// 4
	// Closed!
}
