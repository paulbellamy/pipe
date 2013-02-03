// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

type TakeWhileFunc func(item interface{}) bool

// Accept items from the input pipe until the given function returns false.
// After that, all input messages will be ignored and the output channel will
// be closed.
func TakeWhile(input chan interface{}, fn TakeWhileFunc) chan interface{} {
	output := make(chan interface{})
	go func() {
		for {
			item, ok := <-input
			if !ok {
				break
			}

			// check if we should continue
			if !fn(item) {
				break
			}

			output <- item
		}

		// hit the toggle, close the channel
		close(output)

		// drop any extra messages
		for {
			_, ok := <-input
			if !ok {
				break
			}
		}
	}()
	return output
}

// Helper for the chained constructor
func (p *Pipe) TakeWhile(fn TakeWhileFunc) *Pipe {
	p.Output = TakeWhile(p.Output, fn)
	return p
}
