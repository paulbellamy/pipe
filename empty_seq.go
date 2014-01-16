package pipe

type emptySeq struct{}

var Empty = &emptySeq{}

func (s *emptySeq) First() interface{} {
	return nil
}

func (s *emptySeq) Rest() Seq {
	return s
}

func (s *emptySeq) Empty() bool {
	return true
}
