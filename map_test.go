// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
	"testing"
)

func TestMapChan(t *testing.T) {
	count := 0
	counter := func(item int) string {
		count++
		return fmt.Sprint(count)
	}
	in := make(chan int, 5)
	out := Map(counter, in).(chan string)

	go func() {
		in <- 7
		in <- 4
		in <- 5
	}()
	for i := 1; i <= 3; i++ {
		if result := <-out; result != fmt.Sprint(i) {
			t.Fatal("MapChan received ", i, " items but output ", result)
		}
	}

	close(in)
}

func TestMapSlice(t *testing.T) {
	count := 0
	counter := func(item int) string {
		count++
		return fmt.Sprint(count)
	}
	in := []int{7, 4, 5}
	out := Map(counter, in).([]string)

	for i := 1; i <= 3; i++ {
		if result := out[i-1]; result != fmt.Sprint(i) {
			t.Fatal("MapSlice received", in, "but output", out, "expected", []string{"1", "2", "3"})
		}
	}
}

func TestMapMap(t *testing.T) {
	count := 0
	counter := func(key int, item string) string {
		count++
		return fmt.Sprint(count)
	}
	in := map[int]string{
		7: "a",
		4: "b",
		5: "c",
	}
	out := Map(counter, in).([]string)

	for i := 1; i <= 3; i++ {
		if result := out[i-1]; result != fmt.Sprint(i) {
			t.Fatal("MapSlice received", in, "but output", out, "expected", []string{"1", "2", "3"})
		}
	}
}
