// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Accept only the given number of items from the input pipe. After that number
// has been received, all input messages will be ignored and the output channel
// will be closed.
func Take(input chan interface{}, num int64) chan interface{} {
	output := make(chan interface{})
	var count int64
	go func() {
		// only send num items
		for count = 0; count < num; count++ {
			item, ok := <-input
			if !ok {
				break
			}

			output <- item
		}

		// sent our max, close the channel
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
func (p *Pipe) Take(num int64) *Pipe {
	p.Output = Take(p.Output, num)
	return p
}
