package pipe

// Persistent, functional, sequence interface
type Seq interface {
  First() interface{}
  Rest() Seq
  IsEmpty() bool
}
