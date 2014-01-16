package pipe

import (
	"reflect"
)

type chanSeq struct {
	source reflect.Value
	closed bool
	first  interface{}
}

// Create a new Seq from a chan.
func FromChan(source interface{}) Seq {
	s := &chanSeq{
		source: ensureValue(source),
		closed: false,
	}
	s.first = s
	return s
}

func (s *chanSeq) First() interface{} {
	if s.first != s {
		// If we've already taken a value off the chan
		return s.first
	} else {
		if first, ok := s.source.Recv(); ok {
			s.first = first.Interface()
			return s.first
		} else {
			s.closed = true
			return nil
		}
	}
}

func (s *chanSeq) Rest() Seq {
	s.First()
	if s.closed {
		return s
	} else {
		rest := &chanSeq{
			source: s.source,
			closed: s.closed,
		}
		rest.first = rest
		return rest
	}
}

func (s *chanSeq) Empty() bool {
	s.First() // force lazy eval
	return s.closed
}
