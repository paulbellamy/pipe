// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

type FakeFilter struct {
	count int
}

// Only let even numbers through
func (t *FakeFilter) Filter(item interface{}) bool {
	return (item.(int) % 2) == 0
}

func TestFilterFuncPipe(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	pipe := NewPipe(in, out)
	pipe.FilterFunc(func(item interface{}) bool {
		return (item.(int) % 2) == 0
	})

	in <- 7
	in <- 4
	if result := <-out; result != 4 {
		t.Fatal("even pipe received 7 and 4 but output ", result)
	}

	close(in)
}

func TestFilterPipe(t *testing.T) {
	in := make(chan interface{}, 10)
	out := make(chan interface{}, 10)
	pipe := NewPipe(in, out)
	pipe.Filter(&FakeFilter{})

	// Push in some numbers
	for i := 0; i < 5; i++ {
		in <- i
	}

	// Check only evens came out
	var result interface{}
	for i := 0; i < 5; i += 2 {
		result = <-out
		if result.(int) != i {
			t.Fatal("even object pipe let slip ", result.(int))
		}
	}

	close(in)
}
