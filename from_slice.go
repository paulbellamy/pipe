package pipe

import (
	"reflect"
)

type sliceSeq struct {
	source reflect.Value
}

// Create a new Seq from a slice.
func FromSlice(source interface{}) Seq {
	return &sliceSeq{ensureValue(source)}
}

func (s *sliceSeq) First() interface{} {
	return s.source.Index(0).Interface()
}

func (s *sliceSeq) Rest() Seq {
	return FromSlice(s.source.Slice(1, s.source.Len()))
}

func (s *sliceSeq) Empty() bool {
	return s.source.Len() == 0
}
