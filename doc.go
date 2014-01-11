// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/*
Package pipe provides concurrent and (relatively) transparent transformations
along Golang channels.

For example, to count the number of items passing through a channel:

  // Define our counter
  var counter int
  counter_func := func(item interface{}) {
    counter++ // increment the counter for each item
  }

  // Set up our pipe
  input := make(chan interface{}, 5)

  // Add our counter
  output := ForEach(input, counter_func)

  // Now we send some items
  input <- true
  input <- true
  input <- true

  // Check how many have gone through
  fmt.Println(counter) // prints "3"

You can, of course, modify the items flowing through the pipe:

  // Set up our pipe
  input := make(chan interface{}, 5)

  map_func := func(item interface{}) interface{} {
    // Add 2 to each
    return item.(int) + 2
  }

  filter_func := func(item interface{}) bool {
    // Only allow ints divisible by 5
    return (item.(int) % 5) == 0
  }

  output := Filter(input, filter_func)
  output = Map(output, map_func)

  // Now we send some items
  input <- 1 // will be dropped
  input <- 5 // will come through as 7

There is also a nicer syntax for building sequential pipes:

  // Set up our pipe
  input := make(chan interface{}, 5)

  output := NewPipe(input). // Take items from 'input'
    Filter(func(item interface{}) bool {
      // Only allow items divisible by 5
      return (item.(int) % 5) == 0
    }).
    Map(func(item interface{}) interface{} {
      // Then add 2 to each item
      return item.(int) + 2
    }).
    Output

  // Now we send some items
  input <- 1 // will be dropped
  input <- 5 // will come through as 7
*/
package pipe
