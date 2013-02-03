// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Skip a given number of items from the input pipe. After that number has been
// dropped, the rest are passed straight through.
func Skip(input chan interface{}, num int64) chan interface{} {
	output := make(chan interface{})
	var count int64
	go func() {
		// skip num items
		for count = 0; count < num; count++ {
			_, ok := <-input
			if !ok {
				// channel closed early
				close(output)
				return
			}
		}

		// Return the rest
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

// Helper for chained constructor
func (p *Pipe) Skip(num int64) *Pipe {
	p.Output = Skip(p.Output, num)
	return p
}
