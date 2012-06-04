// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "testing"
)

type CountTransformer struct {
  count int
}

func (t *CountTransformer) Transform(item interface{}) interface{} {
  t.count++
  return t.count
}

func TestNullPipe(t *testing.T) {
  in := make(chan interface{})
  out := make(chan interface{})
  pipe := NewPipe(in, out)

  in <- 5
  if result := <-out; result != 5 {
    t.Fatal("Null pipe received: 5 but output ",result)
  }

  pipe.AddFunc(func(item interface{}) interface{} {
    return item.(int) + 5
  })
}

func TestSinglePipe(t *testing.T) {
  in := make(chan interface{})
  out := make(chan interface{})
  pipe := NewPipe(in, out)
  pipe.AddFunc(func(item interface{}) interface{} {
    return item.(int) + 5
  })

  in <- 5
  if result := <-out; result != 10 {
    t.Fatal("+5 pipe received: 5 but output ",result)
  }
}

func TestMultiPipe(t *testing.T) {
  in := make(chan interface{})
  out := make(chan interface{})
  pipe := NewPipe(in, out)
  pipe.AddFunc(func(item interface{}) interface{} {
    return item.(int) + 5
  })
  pipe.AddFunc(func(item interface{}) interface{} {
    return item.(int) * 2
  })

  in <- 5
  if result := <-out; result != 20 {
    t.Fatal("(x+5)*2 pipe received: 5 but output ",result)
  }
}

func TestObjectPipe(t *testing.T) {
  in := make(chan interface{}, 10)
  out := make(chan interface{}, 10)
  pipe := NewPipe(in, out)
  pipe.Add(&CountTransformer{})

  for i := 0; i < 5; i++ {
    in <- 0
  }
  // find the last item
  var result interface{}
  for i := 0; i < 5; i++ {
    result = <-out
  }

  if result.(int) != 5 {
    t.Fatal("counting pipe received 5 elements but last output was ",result.(int))
  }
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
