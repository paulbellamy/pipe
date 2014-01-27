package pipe

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	printer := func(item int) []string {
		return []string{fmt.Sprintf("%d", item), fmt.Sprintf("%.3d", item)}
	}
	in := make(chan int, 5)
	out := Chain(in).MapCat(printer).Chan().(chan string)

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
			t.Fatal("MapCat received ", i, " items but output ", result)
		}
		if result := <-out; result != fmt.Sprintf("%.3d", i) {
			t.Fatal("MapCat received ", i, " items but output ", result)
		}
	}

	if _, ok := <-out; ok {
		t.Fatal("Expected MapCat to be closed")
	}
}
