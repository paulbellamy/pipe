// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestZipPipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	other := make(chan interface{}, 5)
	NewPipe(in, out).Zip(other)

	in <- 5
	in <- 10
	in <- 20
	other <- 6
	other <- 11

	for i := 1; i <= 2; i++ {
		result := <-out
		expected := []int{i * 5, (i * 5) + 1}
		if len(result.([]interface{})) != len(expected) {
			t.Fatal("expected channel output to match", expected, "but got", result.([]int))
		}

		for j := 0; j < len(result.([]interface{})); j++ {
			if result.([]interface{})[j].(int) != expected[j] {
				t.Fatal("expected channel output to match", expected, "but got", result.([]interface{}))
			}
		}
	}
}
