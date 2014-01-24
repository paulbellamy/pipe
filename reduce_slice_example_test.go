package pipe

import (
	"fmt"
)

func concat(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func ExampleReduceSlice() {
	// Declare a slice of some things

	chars := []string{"a", "b", "c"}

	sum := ReduceSlice(concat, "", chars).(string)

	fmt.Println(sum)

	// Output:
	// abc
}
