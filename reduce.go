// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

type ReduceFunc func(result, item interface{}) interface{}

// Accumulate the result of the reduce function being called on each item, then
// when the input channel is closed, pass the result to the output channel
func Reduce(input chan interface{}, initial interface{}, fn ReduceFunc) chan interface{} {
	output := make(chan interface{})
	var result interface{} = initial
	go func() {
		for {
			item, ok := <-input
			if !ok {
				break
			}

			result = fn(result, item)
		}
		// Input was closed, send the result
		output <- result
		close(output)
	}()
	return output
}

// Accumulate the result of the reduce function being called on each item, then
// when the input channel is closed, pass the result to the output channel
func (p *Pipe) Reduce(initial interface{}, fn ReduceFunc) *Pipe {
	p.Output = Reduce(p.Output, initial, fn)
	return p
}
