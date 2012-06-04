// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Add a transformation to the end of the pipe
func (p *Pipe) Take(num int64) {
  p.addStage()
	go p.takerHandler(num, p.length-1)()
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
