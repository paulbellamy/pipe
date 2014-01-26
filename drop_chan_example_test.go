package pipe

import (
	"fmt"
)

func ExampleDropChan() {
	in := make(chan int, 10)
	out := DropChan(3, in).(chan int)

	for i := 0; i < 5; i++ {
		in <- i
	}

  fmt.Println(<-out)
  fmt.Println(<-out)

  // Output:
  // 3
  // 4
}
