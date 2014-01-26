package pipe

import (
	"fmt"
)

func ExampleFlattenChan() {
	in := make(chan []int, 5)
	out := FlattenChan(in).(chan int)

	in <- []int{1, 2}
	in <- []int{3, 4}
	in <- []int{5, 6}
	close(in)

	for result := range out {
		fmt.Println(result)
	}
	fmt.Println("Closed!")

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// Closed!
}
