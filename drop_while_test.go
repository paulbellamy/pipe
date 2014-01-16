// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestDropWhilePipe(t *testing.T) {
	in := make(chan int, 5)
	out := DropWhile(func(item int) bool {
		return item < 3
	}, in).(chan int)

	in <- 1
	in <- 2
	in <- 3

	result := <-out
	if result != 3 {
		t.Fatal("DropWhile should have dropped all results until 3, but output", result)
	}

	close(in)
}
