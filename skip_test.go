// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestSkipPipe(t *testing.T) {
	in := make(chan interface{}, 10)
	out := make(chan interface{}, 10)
	pipe := NewPipe(in, out)
	pipe.Skip(3)

	for i := 0; i < 5; i++ {
		in <- i
	}

	received := []int{}
	for len(received) < 2 {
		result, ok := <-out
		if !ok {
			break
		}
		received = append(received, result.(int))
	}

	if len(received) != 2 || received[0] != 3 || received[1] != 4 {
		t.Fatal("skip(3) pipe received 1..4 but output ", received)
	}

	close(in)
}
