// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

type FilterFunc func(item interface{}) bool

// Apply a filtering function to a channel, which will only pass through items
// when the filter func returns true.
func Filter(input chan interface{}, fn FilterFunc) chan interface{} {
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

// Helper for chained construction
func (p *Pipe) Filter(fn FilterFunc) *Pipe {
	p.Output = Filter(p.Output, fn)
	return p
}
