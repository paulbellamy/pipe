// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Implement this interface in your object to pass it to Pipe.Add
type Transformer interface {
  Transform(item interface{}) interface{}
}

type transformer func(in interface{}) interface{}

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
	pipe.AddFunc(func(item interface{}) interface{} {
		return item
	})

	return pipe
}

// Add a transformation to the end of the pipe
func (p *Pipe) AddFunc(fn transformer) {
	p.length++
	for i := 0; i < p.length; i++ {
		p.inputs = append(p.inputs, make(chan interface{}))
	}
	go p.handler(fn, p.length-1)()
}

// Add a transformation to the end of the pipe
func (p *Pipe) Add(t Transformer) {
  p.AddFunc(func(item interface{}) interface{} {
    return t.Transform(item)
  })
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

func (p *Pipe) handler(fn transformer, pos int) func() {
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				// channel closed cascade the close
				close(p.nextChan(pos))
				return
			}

			p.nextChan(pos) <- fn(item)
		}
	}
}
