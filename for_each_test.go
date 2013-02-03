// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestForEach(t *testing.T) {
	count := 0

	in := make(chan interface{}, 5)
	out := ForEach(in, func(item interface{}) {
		count++
	})

	in <- 5
	in <- 6
	in <- 7

	// drain the pipe
	for i := 5; i <= 7; i++ {
		result := <-out
		if result.(int) != i {
			t.Fatal("counting ForEach pipe modified ", i, " into ", result.(int))
		}
	}

	if count != 3 {
		t.Fatal("counting ForEach pipe received 3 items but counted ", count)
	}

	close(in)
}

func TestForEachChainedConstructor(t *testing.T) {
	count := 0
	in := make(chan interface{}, 10)
	out := NewPipe(in).
		ForEach(func(item interface{}) {
		count++
	}).
		Output

	// Push in some numbers
	for i := 0; i < 5; i++ {
		in <- i
	}

	// Check it didn't modify
	var result interface{}
	for i := 0; i < 5; i++ {
		result = <-out
		if result.(int) != i {
			t.Fatal("ForEachPipe modified ", i, " into ", result.(int))
		}
	}

	if count != 5 {
		t.Fatal("ForEachPipe miscounted ", 5, " elements as ", count)
	}

	close(in)
}
