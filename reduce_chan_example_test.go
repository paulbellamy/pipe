package pipe

import (
	"fmt"
)

func ExampleReduceChan() {
	// Declare a slice of some things

  chars := make(chan string)

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
