package main

import (
	"fmt"
	. "pipe"
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

func main() {
	places := []place{
		{"Grantchester", 552},
		{"Cambridge", 117900},
		{"Prague", 1188126},
	}
	fmt.Println(Map(statusByPopulation, places))
}
