// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

type FakeTakeWhiler struct {
	Continue bool
}

// returns the index of each element
func (t *FakeTakeWhiler) TakeWhile(item interface{}) bool {
	return t.Continue
}

func TestTakeWhileFuncPipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	pipe := NewPipe(in, out)
	take := true
	pipe.TakeWhileFunc(func(item interface{}) bool {
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

func TestTakeWhilePipe(t *testing.T) {
	in := make(chan interface{}, 10)
	out := make(chan interface{}, 10)
	pipe := NewPipe(in, out)
	taker := &FakeTakeWhiler{}
	pipe.TakeWhile(taker)

	// Push in some numbers
	in <- 1
	in <- 2
	taker.Continue = false // turn off the taker
	in <- 3

	<-out
	<-out
	if _, ok := <-out; ok {
		t.Fatal("takewhile pipe should have closed the channel after turning it off")
	}

	close(in)
}
