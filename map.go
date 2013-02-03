// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

type MapFunc func(item interface{}) interface{}

// Pass through the result of the map function for each item
func Map(input chan interface{}, fn MapFunc) chan interface{} {
	output := make(chan interface{})
	go func() {
		for {
			item, ok := <-input
			if !ok {
				break
			}

			output <- fn(item)
		}
		close(output)
	}()
	return output
}

// Helper for chained construction
func (p *Pipe) Map(fn MapFunc) *Pipe {
	p.Output = Map(p.Output, fn)
	return p
}
