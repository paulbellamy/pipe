// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestReducePipe(t *testing.T) {
	in := make(chan int, 5)
	out := Reduce(in, 0, func(sum, item int) int {
		return sum + item
	}).(chan int)

	in <- 5
	in <- 10
	in <- 20
	close(in)

	result, ok := <-out
	if !ok {
		t.Fatal("output channel was closed before we retrieved the result")
	}

	if result != 35 {
		t.Fatal("reducing (sum) pipe received 5, 10, and 20 items but output ", result)
	}
}
