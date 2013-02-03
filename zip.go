// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Group each message from the input channel with it's corresponding message
// from the other channel. This will block on the first channel until it
// receives a message, then block on the second until it gets one from there.
// At that point an array containing both will be sent to the output channel.
//
// For example, if channel a is being zipped with channel b, and output on channel c:
//
//   a <- 1
//   b <- 2
//   result := <-c // result will equal []interface{}{1, 2}
//
func Zip(input chan interface{}, other chan interface{}) chan interface{} {
	output := make(chan interface{})
	go func() {
		// only send num items
		for {
			a, ok := <-input
			if !ok {
				break
			}

			b, ok := <-other
			if !ok {
				break
			}

			output <- []interface{}{a, b}
		}

		close(output)
	}()
	return output
}

// Helper for the chained constructor
func (p *Pipe) Zip(other chan interface{}) *Pipe {
	p.Output = Zip(p.Output, other)
	return p
}
