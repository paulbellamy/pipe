// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Apply a filtering function to a channel, which will only pass through items
// when the filter func returns true.
func Filter(fn func(item interface{}) bool, input chan interface{}) chan interface{} {
	output := make(chan interface{})
	go func() {
		for {
			item, ok := <-input
			if !ok {
				break
			}

			if fn(item) {
				output <- item
			}
		}
		close(output)
	}()
	return output
}
