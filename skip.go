// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Skip a given number of items from the input pipe. After that number has been
// dropped, the rest are passed straight through.
func (p *Pipe) Skip(num int64) *Pipe {
	p.addStage()
	go p.skipperHandler(num, p.length-1)()

	return p
}

func (p *Pipe) skipperHandler(num int64, pos int) func() {
	var count int64
	return func() {
		// skip num items
		for count = 0; count < num; count++ {
			_, ok := <-p.prevChan(pos)
			if !ok {
				// channel closed early
				close(p.nextChan(pos))
				return
			}
		}

		// Return the rest
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
