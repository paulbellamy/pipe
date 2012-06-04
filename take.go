// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Accept only the given number of items from the input pipe. After that number
// has been received, all input messages will be ignored and the output channel
// will be closed.
func (p *Pipe) Take(num int64) *Pipe {
	p.addStage()
	go p.takerHandler(num, p.length-1)()

	return p
}

func (p *Pipe) takerHandler(num int64, pos int) func() {
	var count int64
	return func() {
		// only send num items
		for count = 0; count < num; count++ {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			p.nextChan(pos) <- item
		}

		// sent our max, close the channel
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
