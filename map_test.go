// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
	"testing"
)

func TestMapPipe(t *testing.T) {
	count := 0
	counter := func(item int) string {
		count++
		return fmt.Sprint(count)
	}
	in := make(chan int, 5)
	out := Map(in, counter).(chan string)

	go func() {
		in <- 7
		in <- 4
		in <- 5
	}()
	for i := 1; i <= 3; i++ {
		if result := <-out; result != fmt.Sprint(i) {
			t.Fatal("mapping pipe received ", i, " items but output ", result)
		}
	}

	close(in)
}
