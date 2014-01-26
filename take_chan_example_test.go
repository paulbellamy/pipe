package pipe

import (
	"fmt"
)

func ExampleTakeChan() {
	in := make(chan int, 10)
	out := TakeChan(3, in).(chan int)

	for i := 0; i < 5; i++ {
		in <- i
	}

	for result := range out {
		fmt.Println(result)
	}
	fmt.Println("Closed!")

	// Output:
	// 0
	// 1
	// 2
	// Closed!
}
