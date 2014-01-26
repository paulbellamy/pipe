package pipe

import (
	"fmt"
	"testing"
)

func TestMapCatChan(t *testing.T) {
	printer := func(item int) []string {
		return []string{fmt.Sprintf("%d", item), fmt.Sprintf("%.3d", item)}
	}
	in := make(chan int, 5)
	out := MapCatChan(printer, in).(chan string)

	go func() {
		in <- 1
		in <- 2
		in <- 3
		close(in)
	}()

	count := 0
	for i := 1; i <= 3; i++ {
		count++
		if result := <-out; result != fmt.Sprintf("%d", i) {
			t.Fatal("MapCatChan received ", i, " items but output ", result)
		}
		if result := <-out; result != fmt.Sprintf("%.3d", i) {
			t.Fatal("MapCatChan received ", i, " items but output ", result)
		}
	}

	if _, ok := <-out; ok {
		t.Fatal("Expected MapCatChan to be closed")
	}
}
