// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

type FakeSkipWhiler struct {
}

// returns the index of each element
func (t *FakeSkipWhiler) SkipWhile(item interface{}) bool {
	return item.(int) < 3
}

func TestSkipWhileFuncPipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	NewPipe(in, out).SkipWhileFunc(func(item interface{}) bool {
		return item.(int) < 3
	})

	in <- 1
	in <- 2
	in <- 3

	result := <-out
	if result != 3 {
		t.Fatal("skipwhile should have skipped all results until 3, but output", result)
	}

	close(in)
}

func TestSkipWhilePipe(t *testing.T) {
	in := make(chan interface{}, 10)
	out := make(chan interface{}, 10)
	skipper := &FakeSkipWhiler{}
	NewPipe(in, out).SkipWhile(skipper)

	// Push in some numbers
	in <- 1
	in <- 2
	in <- 3

	result := <-out
	if result != 3 {
		t.Fatal("skipwhile should have skipped all results until 3, but output", result)
	}

	close(in)
}
