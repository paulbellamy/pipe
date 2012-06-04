// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Implement this interface in your object to pass it to Pipe.SkipWhile
type SkipWhiler interface {
	SkipWhile(item interface{}) bool
}

// A function which skipwhiles
type SkipWhileFunc func(item interface{}) bool

// Add a transformation to the end of the pipe
func (p *Pipe) SkipWhileFunc(fn SkipWhileFunc) *Pipe {
	p.addStage()
	go p.skipwhileHandler(fn, p.length-1)()

	return p
}

// Add a transformation to the end of the pipe
func (p *Pipe) SkipWhile(t SkipWhiler) *Pipe {
	p.SkipWhileFunc(func(item interface{}) bool {
		return t.SkipWhile(item)
	})

	return p
}

func (p *Pipe) skipwhileHandler(fn SkipWhileFunc, pos int) func() {
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				// input closed, abort
				close(p.nextChan(pos))
				return
			}

			// check if we should output this
			if !fn(item) {
				p.nextChan(pos) <- item
				break
			}
		}

		// send any messages after this
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			p.nextChan(pos) <- item
		}

		close(p.nextChan(pos))

	}
}
