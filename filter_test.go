// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestFilter(t *testing.T) {
	mod2 := func(item interface{}) bool {
		return (item.(int) % 2) == 0
	}

	in := make(chan interface{})
	out := Filter(in, mod2)

	in <- 7
	in <- 4
	if result := <-out; result != 4 {
		t.Fatal("even pipe received 7 and 4 but output ", result)
	}

	close(in)
}

func TestFilterChainedConstructor(t *testing.T) {
	mod2 := func(item interface{}) bool {
		return (item.(int) % 2) == 0
	}

	in := make(chan interface{})
	out := NewPipe(in).Filter(mod2).Output

	in <- 7
	in <- 4
	if result := <-out; result != 4 {
		t.Fatal("even pipe received 7 and 4 but output ", result)
	}

	close(in)
}
