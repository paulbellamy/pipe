// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestTakeWhileFuncPipe(t *testing.T) {
	take := true
	in := make(chan interface{}, 5)
	out := TakeWhile(in, func(item interface{}) bool {
		return take
	})

	in <- 7
	in <- 4
	take = false
	in <- 5

	<-out
	<-out
	if _, ok := <-out; ok {
		t.Fatal("takewhile pipe should have closed the channel after turning it off")
	}

	close(in)
}

func TestTakeWhileChainedConstructor(t *testing.T) {
	take := true
	in := make(chan interface{}, 10)
	out := NewPipe(in).
		TakeWhile(func(item interface{}) bool {
		return take
	}).
		Output

	// Push in some numbers
	in <- 1
	in <- 2
	take = false
	in <- 3

	<-out
	<-out
	if _, ok := <-out; ok {
		t.Fatal("takewhile pipe should have closed the channel after turning it off")
	}

	close(in)
}
