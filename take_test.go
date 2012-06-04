// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestTakePipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	pipe := NewPipe(in, out)
	pipe.Take(3)

	for i := 0; i < 5; i++ {
		in <- i
	}

	count := 0
	for {
		_, ok := <-out
		if !ok {
			break
		}
		count++
	}

	if count != 3 {
		t.Fatal("take(3) pipe received 5 items but output ", count)
	}
}
