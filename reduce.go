// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Implement this interface in your object to pass it to Pipe.Reduce
type Reducer interface {
	Reduce(item interface{}) interface{}
}

// A function which reduces
type ReduceFunc func(result, item interface{}) interface{}

// Add a transformation to the end of the pipe
func (p *Pipe) ReduceFunc(initial interface{}, fn ReduceFunc) *Pipe {
	p.addStage()
	go p.reducerHandler(initial, fn, p.length-1)()

	return p
}

// Add a transformation to the end of the pipe
func (p *Pipe) Reduce(t Reducer) *Pipe {
	p.addStage()
	var pos int = p.length - 1
	var result interface{}
	go func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			result = t.Reduce(item)
		}
		// Input was closed, send the result
		p.nextChan(pos) <- result
		close(p.nextChan(pos))
	}()

	return p
}

func (p *Pipe) reducerHandler(initial interface{}, fn ReduceFunc, pos int) func() {
	var result interface{} = initial
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			result = fn(result, item)
		}
		// Input was closed, send the result
		p.nextChan(pos) <- result
		close(p.nextChan(pos))
	}
}
