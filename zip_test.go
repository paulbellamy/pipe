// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestZipPipe(t *testing.T) {
	other := make(chan int, 5)
	in := make(chan int, 5)
	out := Zip(in, other).(chan []int)

	in <- 5
	in <- 10
	in <- 20
	other <- 6
	other <- 11

	for i := 1; i <= 2; i++ {
		result := <-out
		expected := []int{i * 5, (i * 5) + 1}
		if len(result) != len(expected) {
			t.Fatal("expected channel output to match", expected, "but got", result)
		}

		for j := 0; j < len(result); j++ {
			if result[j] != expected[j] {
				t.Fatal("expected channel output to match", expected, "but got", result)
			}
		}
	}
}
