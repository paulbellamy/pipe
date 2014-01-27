package pipe

type emptySeq struct {}

func (s emptySeq) First() interface{} {
  return nil
}

func (s emptySeq) Rest() Seq {
  return s
}

func (s emptySeq) IsEmpty() bool {
  return true
}

var Empty Seq = emptySeq{}
