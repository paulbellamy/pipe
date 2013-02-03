// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

type SkipWhileFunc func(item interface{}) bool

// Skip the items from the input pipe until the given function returns true.
// After that , the rest are passed straight through.
func SkipWhile(input chan interface{}, fn SkipWhileFunc) chan interface{} {
	output := make(chan interface{})
	go func() {
		for {
			item, ok := <-input
			if !ok {
				// input closed, abort
				close(output)
				return
			}

			// check if we should output this
			if !fn(item) {
				output <- item
				break
			}
		}

		// send any messages after this
		for {
			item, ok := <-input
			if !ok {
				break
			}

			output <- item
		}

		close(output)

	}()
	return output
}

// Helper function for chained constructor
func (p *Pipe) SkipWhile(fn SkipWhileFunc) *Pipe {
	p.Output = SkipWhile(p.Output, fn)
	return p
}
