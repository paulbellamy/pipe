package pipe

import (
	"fmt"
	"testing"
)

func TestReduceChan(t *testing.T) {
	in := make(chan int, 5)

	go func() {
		in <- 5
		in <- 10
		in <- 20
		close(in)
	}()

	out := ReduceChan(sum, 0, in).(int)

	if out != 35 {
		t.Fatal("ReduceChan(sum, 0, []int{5, 10, 20}) output ", out)
	}
}

func TestReduceChanTypeCoercion(t *testing.T) {
	appendToString := func(str string, item fmt.Stringer) string {
		return fmt.Sprintf("%s%s", str, item.String())
	}

	in := make(chan testStringer, 5)

	go func() {
		in <- 1
		in <- 2
		in <- 3
		close(in)
	}()

	out := ReduceChan(appendToString, "a", in).(string)

	if out != "a123" {
		t.Fatal("ReduceChan(appendToString, \"a\", chan testStringer) output ", out)
	}
}
