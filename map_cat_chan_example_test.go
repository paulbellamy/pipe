package pipe

import (
	"fmt"
)

func ExampleMapCatChan() {
	printer := func(item int) []string {
		return []string{fmt.Sprintf("%d", item), fmt.Sprintf("%.3d", item)}
	}
	in := make(chan int, 5)
	out := MapCatChan(printer, in).(chan string)

	go func() {
		in <- 1
		in <- 2
		in <- 3
		close(in)
	}()

	for result := range out {
		fmt.Println(result)
	}
	fmt.Println("Closed!")

	// Output
	// 1
	// 001
	// 2
	// 002
	// 3
	// 003
	// Closed!
}
