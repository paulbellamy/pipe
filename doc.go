// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/*
Package pipe provides concurrent and (relatively) transparent transformations
along Golang channels.

For example, to count the number of items passing through a channel:

  // Define our counter
  type PipeCounter struct {
    Count int
  }

  // tell it what to do with each item
  func (c *PipeCounter) ForEach(item interface{}) {
    c.Count++ // increment the counter
  }

  // Set up our pipe
  input := make(chan interface{}, 5)
  output := make(chan interface{}, 5)
  pipe := NewPipe(input, output)

  // Add our counter
  counter := &PipeCounter{}
  pipe.ForEach(counter)

  // Now we send some items
  input <- true
  input <- true
  input <- true

  // Check how many have gone through
  fmt.Println(counter.Count) // prints "3"

You can, of course, modify the items flowing through the pipe:

  // Set up our pipe
  input := make(chan interface{}, 5)
  output := make(chan interface{}, 5)
  pipe := NewPipe(input, output)

  // Only allow ints divisible by 5
  pipe.Filter(func(item interface{}) bool {
    return (item.(int) % 5) == 0
  })

  // Now we send some items
  input <- 1 // will be dropped
  input <- 5 // will come through
*/
package pipe
