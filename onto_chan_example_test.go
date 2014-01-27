package pipe

import (
	"fmt"
)

func ExampleOntoChan() {
	in := []int{0, 1, 2, 3, 4}
	out := OntoChan(in).(chan int)

	for x := range out {
		fmt.Println(x)
	}
	fmt.Println("Closed!")

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// Closed!
}
