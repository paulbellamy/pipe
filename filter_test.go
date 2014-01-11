// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestFilter(t *testing.T) {
	mod2 := func(item int) bool {
		return (item % 2) == 0
	}

	in := make(chan int)
	out := Filter(Filter(in, mod2), mod2).(chan int)

	in <- 7
	in <- 4
	if result := <-out; result != 4 {
		t.Fatal("even pipe received 7 and 4 but output ", result)
	}

	close(in)
}
