// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestTakeWhileFuncPipe(t *testing.T) {
	take := true
	in := make(chan int, 5)
	out := TakeWhile(in, func(item int) bool {
		return take
	}).(chan int)

	in <- 7
	in <- 4
	take = false
	in <- 5

	<-out
	<-out
	if _, ok := <-out; ok {
		t.Fatal("takewhile pipe should have closed the channel after turning it off")
	}

	close(in)
}
