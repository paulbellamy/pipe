package pipe

import (
	"testing"
)

func TestDropChan(t *testing.T) {
	in := make(chan int, 10)
	out := DropChan(3, in).(chan int)

	for i := 0; i < 5; i++ {
		in <- i
	}

	received := []int{}
	for len(received) < 2 {
		result, ok := <-out
		if !ok {
			break
		}
		received = append(received, result)
	}

	if len(received) != 2 || received[0] != 3 || received[1] != 4 {
		t.Fatal("DropChan(3) pipe received 1..4 but output ", received)
	}

	close(in)
}
