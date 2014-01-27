package pipe

import (
	"fmt"
	"testing"
)

func TestIterateChan(t *testing.T) {
	fib := func(f, s int) (int, int) {
		return s, f + s
	}

	out := IterateChan(fib, 0, 1).(chan int)

	assertEqual(t, <-out, 0)
	assertEqual(t, <-out, 1)
	assertEqual(t, <-out, 1)
	assertEqual(t, <-out, 2)
	assertEqual(t, <-out, 3)
	assertEqual(t, <-out, 5)
	assertEqual(t, <-out, 8)
	assertEqual(t, <-out, 13)
}

func TestIterateChanTypeCoercion(t *testing.T) {
	counter := func(output fmt.Stringer, state int) (testStringer, int) {
		return testStringer(state), state + 1
	}
	out := IterateChan(counter, testStringer(0), 1).(chan testStringer)

	for i := 0; i <= 3; i++ {
		if result := <-out; result.String() != fmt.Sprint(i) {
			t.Fatal("IterateChan output ", result, "but was expected to output", i)
		}
	}
}
