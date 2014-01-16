// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestDropChan(t *testing.T) {
	in := make(chan int, 10)
	out := Drop(3, in).(chan int)

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
		t.Fatal("Drop(3) pipe received 1..4 but output ", received)
	}

	close(in)
}

func TestDropSlice(t *testing.T) {
	in := []int{0, 1, 2, 3, 4}
	out := Drop(3, in).([]int)

	if len(out) != 2 || out[0] != 3 || out[1] != 4 {
		t.Fatal("Drop(3) expected", []int{3, 4}, "but output ", out)
	}
}

func TestDropSliceNone(t *testing.T) {
	in := []int{0, 1, 2, 3, 4}
	out := Drop(0, in).([]int)

	if len(out) != 5 {
		t.Fatal("Drop(3) expected", in, "but output ", out)
	}
}

func TestDropSliceEmpty(t *testing.T) {
	in := []int{}
	out := Drop(3, in).([]int)

	if len(out) != 0 {
		t.Fatal("Drop(3) expected", in, "but output ", out)
	}
}

func TestDropSliceMoreThanRemains(t *testing.T) {
	in := []int{0, 1, 2, 3, 4}
	out := Drop(6, in).([]int)

	if len(out) != 0 {
		t.Fatal("Drop(3) expected", []int{}, "but output ", out)
	}
}
