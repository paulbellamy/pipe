package pipe

import (
  "fmt"
)

func ExampleDropWhileChan() {
	in := make(chan int, 5)

  lessThan := func(x int) (func (int) bool) {
    return func(item int) bool {
      return item < 3
    }
  }

	out := DropWhileChan(lessThan(3), in).(chan int)

	in <- 1
	in <- 2
	in <- 3
  in <- 2

  fmt.Println(<-out)
  fmt.Println(<-out)

  // Output:
  // 3
  // 2
}
