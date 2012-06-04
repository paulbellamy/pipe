// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

type FakeReducer struct {
	sum int
}

// sum the elements
func (t *FakeReducer) Reduce(item interface{}) interface{} {
	t.sum += item.(int)
	return t.sum
}

func TestReduceFuncPipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	NewPipe(in, out).ReduceFunc(0, func(sum, item interface{}) interface{} {
		return sum.(int) + item.(int)
	})

	in <- 5
	in <- 10
	in <- 20
	close(in)

	result, ok := <-out
	if !ok {
		t.Fatal("output channel was closed before we retrieved the result")
	}

	if result.(int) != 35 {
		t.Fatal("reducing (sum) pipe received 5, 10, and 20 items but output ", result.(int))
	}
}

func TestReducePipe(t *testing.T) {
	in := make(chan interface{}, 10)
	out := make(chan interface{}, 10)
	NewPipe(in, out).Reduce(&FakeReducer{})

	// Push in some numbers
	for i := 5; i > 0; i-- {
		in <- i
	}

	close(in)

	result, ok := <-out
	if !ok {
		t.Fatal("output channel was closed before we retrieved the result")
	}

	expected := 5 + 4 + 3 + 2 + 1
	if result.(int) != expected {
		t.Fatal("reducing (sum) pipe received 5..1 items but output ", result.(int))
	}
}
