package pipe

import (
	"fmt"
)

type place struct {
	name       string
	population int
}

func statusByPopulation(p place) string {
	switch {
	case p.population > 1000000:
		return "City"
	case p.population > 5000:
		return "Town"
	default:
		return "Village"
	}
}

func ExampleMapSlice() {
	// Declare a slice of some things
	places := []place{
		{"Grantchester", 552},
		{"Cambridge", 117900},
		{"Prague", 1188126},
	}

	sizes := MapSlice(statusByPopulation, places).([]string)

	fmt.Println(sizes)

	// Output:
	// [Village Town City]
}
