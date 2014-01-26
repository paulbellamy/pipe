package pipe

import (
	"testing"
)

func TestTakeChan(t *testing.T) {
	in := make(chan int, 10)
	out := TakeChan(3, in).(chan int)

	for i := 0; i < 5; i++ {
		in <- i
	}

	count := 0
	for {
		_, ok := <-out
		if !ok {
			break
		}
		count++
	}

	if count != 3 {
		t.Fatal("TakeChan(3) received 5 items but output ", count)
	}
}
