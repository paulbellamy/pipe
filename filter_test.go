// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestFilterPipe(t *testing.T) {
	in := make(chan interface{})
  out := Filter(func(item interface{}) bool {
		return (item.(int) % 2) == 0
	}, in)

	in <- 7
	in <- 4
	if result := <-out; result != 4 {
		t.Fatal("even pipe received 7 and 4 but output ", result)
	}

	close(in)
}
