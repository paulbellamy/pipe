// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Implement this interface in your object to pass it to Pipe.Map
type Mapper interface {
	Map(item interface{}) interface{}
}

// A function which mappers
type MapFunc func(item interface{}) interface{}

// Pass through the result of the map function for each item
func (p *Pipe) MapFunc(fn MapFunc) *Pipe {
	p.addStage()
	go p.mapperHandler(fn, p.length-1)()

	return p
}

// Pass through the result of the map function for each item
func (p *Pipe) Map(t Mapper) *Pipe {
	p.MapFunc(func(item interface{}) interface{} {
		return t.Map(item)
	})

	return p
}

func (p *Pipe) mapperHandler(fn MapFunc, pos int) func() {
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			p.nextChan(pos) <- fn(item)
		}
		close(p.nextChan(pos))
	}
}
