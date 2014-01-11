// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestSkipWhilePipe(t *testing.T) {
	in := make(chan int, 5)
	out := SkipWhile(in, func(item int) bool {
		return item < 3
	}).(chan int)

	in <- 1
	in <- 2
	in <- 3

	result := <-out
	if result != 3 {
		t.Fatal("skipwhile should have skipped all results until 3, but output", result)
	}

	close(in)
}
