// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestSkipPipe(t *testing.T) {
	in := make(chan interface{}, 10)
	out := Skip(in, 3)

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

func TestSkipChainedConstructor(t *testing.T) {
	in := make(chan interface{}, 10)
	out := NewPipe(in).Skip(3).Output

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
