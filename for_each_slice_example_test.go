package pipe

import (
	"fmt"
)

func ExampleForEachSlice() {
  // Declare a slice of some things
	places := []string{"Grantchester", "Cambridge", "Prague"}

	_ = ForEachSlice(fmt.Println, places).([]string)

	// Output:
	// Grantchester
  // Cambridge
  // Prague
}
