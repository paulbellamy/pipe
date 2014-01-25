package pipe

import (
	"fmt"
)

func ExampleForEachChan() {
  // Declare a chan of some things
	places := make(chan string, 3)

  processed := ForEachChan(fmt.Println, places).(chan string)

  places <- "Grantchester"
  places <- "Cambridge"
  places <- "Prague"

  <-processed
  <-processed
  <-processed

	// Output:
	// Grantchester
  // Cambridge
  // Prague
}
