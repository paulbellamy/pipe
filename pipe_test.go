// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

type Counter struct {
	count int
}

// returns the index of each element
func (t *Counter) Map(item interface{}) interface{} {
	t.count++
	return t.count
}

// returns the last t.count when the pipe closes
func (t *Counter) Reduce(item interface{}) interface{} {
	t.count++
	return t.count
}

func TestNullPipe(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	NewPipe(in, out)

	in <- 5
	if result := <-out; result != 5 {
		t.Fatal("Null pipe received: 5 but output ", result)
	}

	close(in)
}

func TestMultiPipe(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	NewPipe(in, out).FilterFunc(func(item interface{}) bool {
		return (item.(int) % 5) == 0
	}).FilterFunc(func(item interface{}) bool {
		return (item.(int) % 2) == 0
	})

	in <- 2
	in <- 5
	in <- 10
	if result := <-out; result != 10 {
		t.Fatal("mod 2 and mod 5 pipe received 2, 5 and 10 but output ", result)
	}

	close(in)
}

func TestClosingPipe(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	NewPipe(in, out)

	close(in)
	if _, ok := <-out; ok {
		t.Fatal("closing the input pipe did not cascade to output")
	}
}
