package pipe

import (
	"fmt"
)

func ExampleReduceChan() {
	// Declare a slice of some things
	chars := make(chan string)

	// Function to apply
	concat := func(a, b string) string {
		return fmt.Sprintf("%s%s", a, b)
	}

	// Push some values into the input
	go func() {
		chars <- "a"
		chars <- "b"
		chars <- "c"
		close(chars)
	}()

	sum := ReduceChan(concat, "", chars).(string)

	fmt.Println(sum)

	// Output:
	// abc
}
