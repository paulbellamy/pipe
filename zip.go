// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Add a transformation to the end of the pipe
func (p *Pipe) Zip(other chan interface{}) {
	p.addStage()
	go p.zipperHandler(other, p.length-1)()
}

func (p *Pipe) zipperHandler(other chan interface{}, pos int) func() {
	return func() {
		// only send num items
		for {
			a, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			b, ok := <-other
			if !ok {
				break
			}

			p.nextChan(pos) <- []interface{}{a, b}
		}

		close(p.nextChan(pos))
	}
}
