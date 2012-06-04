// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Implement this interface in your object to pass it to Pipe.ForEach
type ForEacher interface {
	ForEach(item interface{})
}

// A function which foreachs
type ForEachFunc func(item interface{})

// Add a transformation to the end of the pipe
func (p *Pipe) ForEachFunc(fn ForEachFunc) {
	p.addStage()
	go p.foreachHandler(fn, p.length-1)()
}

// Add a transformation to the end of the pipe
func (p *Pipe) ForEach(t ForEacher) {
	p.ForEachFunc(func(item interface{}) {
		t.ForEach(item)
	})
}

func (p *Pipe) foreachHandler(fn ForEachFunc, pos int) func() {
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			fn(item)
			p.nextChan(pos) <- item
		}
		close(p.nextChan(pos))
	}
}
