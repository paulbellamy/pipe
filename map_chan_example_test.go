package pipe

import (
	"fmt"
)

func ExampleMapChan() {
	// Declare a chan of some things
	numbers := make(chan int)

  // function to apply
  square := func(x int) int {
    return x * x
  }

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
