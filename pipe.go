// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A Pipe is a set of transforms being applied along the channel
type Pipe struct {
	length     int
	inputs     []chan interface{}
	output     chan interface{}
	output_cap int
}

// Return a new Pipe object which echoes input to output
func NewPipe(in, out chan interface{}) *Pipe {
	pipe := &Pipe{
		inputs:     []chan interface{}{in},
		output:     out,
		output_cap: cap(out),
	}

	// Add the null handler (just echoes in to output)
	pipe.FilterFunc(func(item interface{}) bool {
		return true
	})

	return pipe
}

// Create a new channel
func (p *Pipe) addStage() (chan interface{}) {
	p.length++
  c := make(chan interface{})
  p.inputs = append(p.inputs, c)
  return c
}

func (p *Pipe) prevChan(pos int) chan interface{} {
	return p.inputs[pos]
}

func (p *Pipe) nextChan(pos int) chan interface{} {
	if pos == (p.length - 1) {
		return p.output
	}
	return p.inputs[pos+1]
}
