package pipe

import (
	"fmt"
)

func ExampleForEachChan() {
	// Declare a chan of some things
	places := make(chan string, 5)

	places <- "Grantchester"
	places <- "Cambridge"
	places <- "Prague"
	close(places)

	ForEachChan(fmt.Println, places)

	// Output:
	// Grantchester
	// Cambridge
	// Prague
}
