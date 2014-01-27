package pipe

import (
	"fmt"
	"math/rand"
)

func ExampleCall() {
	rand.Seed(1)

	type generator func(int) int
	fns := []generator{rand.Intn, rand.Intn, rand.Intn}

	nums := MapCatSlice(Call(100), fns).([]interface{})

	for _, num := range nums {
		fmt.Println(num.(int))
	}

	// Output:
	// 81
	// 87
	// 47
}
