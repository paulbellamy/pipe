package pipe

import (
	"fmt"
)

func ExampleForEachSlice() {
	// Declare a slice of some things
	places := []string{"Grantchester", "Cambridge", "Prague"}

	ForEachSlice(fmt.Println, places)

	// Output:
	// Grantchester
	// Cambridge
	// Prague
}
