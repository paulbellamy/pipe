// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "testing"
)

type Counter struct {
  count int
}

// Only let even numbers through
func (t *Counter) Filter(item interface{}) bool {
  return (item.(int) % 2) == 0
}

// Counts each item as it goes through
func (t *Counter) ForEach(item interface{}) {
  t.count++
}

// returns the index of each element
func (t *Counter) Map(item interface{}) interface{} {
  t.count++
  return t.count
}

// returns the last t.count when the pipe closes
func (t *Counter) Reduce(item interface{}) interface{} {
  t.count++
  return t.count
}

func TestNullPipe(t *testing.T) {
  in := make(chan interface{})
  out := make(chan interface{})
  NewPipe(in, out)

  in <- 5
  if result := <-out; result != 5 {
    t.Fatal("Null pipe received: 5 but output ",result)
  }

  close(in)
}

func TestFilterPipe(t *testing.T) {
  in := make(chan interface{})
  out := make(chan interface{})
  pipe := NewPipe(in, out)
  pipe.FilterFunc(func(item interface{}) bool  {
    return (item.(int) % 2) == 0
  })

  in <- 7
  in <- 4
  if result := <-out; result != 4 {
    t.Fatal("even pipe received 7 and 4 but output ",result)
  }

  close(in)
}

func TestMultiPipe(t *testing.T) {
  in := make(chan interface{})
  out := make(chan interface{})
  pipe := NewPipe(in, out)
  pipe.FilterFunc(func(item interface{}) bool {
    return (item.(int) % 5) == 0
  })
  pipe.FilterFunc(func(item interface{}) bool {
    return (item.(int) % 2) == 0
  })

  in <- 2
  in <- 5
  in <- 10
  if result := <-out; result != 10 {
    t.Fatal("mod 2 and mod 5 pipe received 2, 5 and 10 but output ",result)
  }

  close(in)
}

func TestObjectFilterPipe(t *testing.T) {
  in := make(chan interface{}, 10)
  out := make(chan interface{}, 10)
  pipe := NewPipe(in, out)
  pipe.Filter(&Counter{})

  // Push in some numbers
  for i := 0; i < 5; i++ {
    in <- i
  }

  // Check only evens came out
  var result interface{}
  for i := 0; i < 5; i += 2 {
    result = <-out
    if result.(int) != i {
      t.Fatal("even object pipe let slip ",result.(int))
    }
  }

  close(in)
}

func TestClosingPipe(t *testing.T) {
  in := make(chan interface{})
  out := make(chan interface{})
  NewPipe(in, out)

  close(in)
  if _, ok := <-out; ok {
    t.Fatal("closing the input pipe did not cascade to output")
  }
}
