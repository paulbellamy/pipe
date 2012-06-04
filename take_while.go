// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Implement this interface in your object to pass it to Pipe.TakeWhile
type TakeWhiler interface {
	TakeWhile(item interface{}) bool
}

// A function which takewhiles
type TakeWhileFunc func(item interface{}) bool

// Add a transformation to the end of the pipe
func (p *Pipe) TakeWhileFunc(fn TakeWhileFunc) *Pipe {
	p.addStage()
	go p.takewhileHandler(fn, p.length-1)()

	return p
}

// Add a transformation to the end of the pipe
func (p *Pipe) TakeWhile(t TakeWhiler) *Pipe {
	p.TakeWhileFunc(func(item interface{}) bool {
		return t.TakeWhile(item)
	})

	return p
}

func (p *Pipe) takewhileHandler(fn TakeWhileFunc, pos int) func() {
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			// check if we should continue
			if !fn(item) {
				break
			}

			p.nextChan(pos) <- item
		}

		// hit the toggle, close the channel
		close(p.nextChan(pos))

		// drop any extra messages
		for {
			_, ok := <-p.prevChan(pos)
			if !ok {
				break
			}
		}
	}
}
