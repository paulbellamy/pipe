package pipe

import (
  "fmt"
  "testing"
)

func TestMapChan(t *testing.T) {
	count := 0
	counter := func(item int) string {
		count++
		return fmt.Sprint(count)
	}
	in := make(chan int, 5)
	out := MapChan(counter, in).(chan string)

	go func() {
		in <- 7
		in <- 4
		in <- 5
	}()
	for i := 1; i <= 3; i++ {
		if result := <-out; result != fmt.Sprint(i) {
			t.Fatal("MapChan received ", i, " items but output ", result)
		}
	}

	close(in)
}

func TestMapChanTypeCoercion(t *testing.T) {
	count := 0
	counter := func(item fmt.Stringer) string {
		count++
		return fmt.Sprint(count)
	}
	in := make(chan testStringer, 5)
	out := MapChan(counter, in).(chan string)

	go func() {
		in <- 7
		in <- 4
		in <- 5
	}()

	for i := 1; i <= 3; i++ {
		if result := <-out; result != fmt.Sprint(i) {
			t.Fatal("MapChan received ", i, " items but output ", result)
		}
	}

	close(in)
}
