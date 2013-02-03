// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A function which foreachs
type ForEachFunc func(item interface{})

// Execute a function for each item (without modifying the item). Useful for
// monitoring, logging, or causing some side-effect.
func ForEach(input chan interface{}, fn ForEachFunc) chan interface{} {
	output := make(chan interface{})
	go func() {
		for {
			item, ok := <-input
			if !ok {
				break
			}

			fn(item)
			output <- item
		}
		close(output)
	}()
	return output
}

// Execute a function for each item (without modifying the item)
func (p *Pipe) ForEach(fn ForEachFunc) *Pipe {
	p.Output = ForEach(p.Output, fn)
	return p
}
