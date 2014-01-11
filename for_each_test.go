// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestForEach(t *testing.T) {
	count := 0

	in := make(chan int, 5)
  counter := func(item int) {
		count++
	}
	out := ForEach(ForEach(in, counter), counter).(chan int)

	in <- 5
	in <- 6
	in <- 7

	// drain the pipe
	for i := 5; i <= 7; i++ {
		result := <-out
		if result != i {
			t.Fatal("counting ForEach pipe modified ", i, " into ", result)
		}
	}

	if count != 6 {
		t.Fatal("counting ForEach pipe received 3 items but counted ", count, "/ 2")
	}

	close(in)
}
