package pipe

import (
	"fmt"
)

func ExampleReduceSlice() {
	// Declare a slice of some things
	chars := []string{"a", "b", "c"}

	// Function to apply
	concat := func(a, b string) string {
		return fmt.Sprintf("%s%s", a, b)
	}

	sum := ReduceSlice(concat, "", chars).(string)

	fmt.Println(sum)

	// Output:
	// abc
}
