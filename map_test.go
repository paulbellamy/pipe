// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

type FakeMapper struct {
	count int
}

// returns the index of each element
func (t *FakeMapper) Map(item interface{}) interface{} {
	t.count++
	return t.count
}

func TestMapFuncPipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	pipe := NewPipe(in, out)
	count := 0
	pipe.MapFunc(func(item interface{}) interface{} {
		count++
		return count
	})

	in <- 7
	in <- 4
	in <- 5
	for i := 1; i <= 3; i++ {
		if result := <-out; result.(int) != i {
			t.Fatal("mapping pipe received ", i, " items but output ", result.(int))
		}
	}

	close(in)
}

func TestMapPipe(t *testing.T) {
	in := make(chan interface{}, 10)
	out := make(chan interface{}, 10)
	pipe := NewPipe(in, out)
	pipe.Map(&FakeMapper{})

	// Push in some numbers
	for i := 5; i > 0; i-- {
		in <- i
	}

	// Check their index came out instead
	var result interface{}
	for i := 1; i <= 5; i++ {
		result = <-out
		if result.(int) != i {
			t.Fatal("mapping pipe should have output", i, "but output", result.(int))
		}
	}

	close(in)
}
