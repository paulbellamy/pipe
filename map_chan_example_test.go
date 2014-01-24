package pipe

import (
	"fmt"
)

// Input and output needn't be the same type.
func square(x int) int {
	return x * x
}

func ExampleMapChan() {
	// Declare a chan of some things
	numbers := make(chan int)

	// Create a new chan which applies a function to the original
	squares := MapChan(square, numbers).(chan int)

	// Put some numbers in the original chan
	go func() {
		numbers <- 1
		numbers <- 2
		numbers <- 3
	}()

	fmt.Println(<-squares)
	fmt.Println(<-squares)
	fmt.Println(<-squares)

	// Output:
	// 1
	// 4
	// 9
}
