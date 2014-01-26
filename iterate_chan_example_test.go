package pipe

import (
	"fmt"
)

func ExampleIterateChan() {
	fib := func(f, s int) (int, int) {
		return s, f + s
	}

	out := IterateChan(fib, 0, 1).(chan int)

	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}

	// Output:
	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
}
