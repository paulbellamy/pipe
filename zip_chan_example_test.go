package pipe

import (
	"fmt"
)

func ExampleZipChan() {
	other := make(chan int, 5)
	in := make(chan int, 5)
	out := ZipChan(in, other).(chan []int)

	in <- 5
	in <- 10
	in <- 20
	other <- 6
	other <- 11
	close(other)

	for result := range out {
		fmt.Println(result)
	}
	fmt.Println("Closed!")

	// Output:
	// [5 6]
	// [10 11]
	// Closed!
}
