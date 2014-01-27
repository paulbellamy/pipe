package pipe

import (
	"reflect"
)

// Persistent, functional, sequence interface
type Seq interface {
	First() interface{}
	Rest() Seq
	Empty() bool
}

var seqType = reflect.TypeOf((*Seq)(nil)).Elem()

// Create a new Seq. Supports Chans and Slices.
func New(source interface{}) Seq {
	sourceValue := ensureValue(source)
	if sourceValue.Type().Implements(seqType) {
		var s Seq
		reflect.ValueOf(s).Set(sourceValue)
		return s
	}

	switch sourceValue.Type().Kind() {
	case reflect.Chan:
		return FromChan(sourceValue)
	case reflect.Slice:
		return FromSlice(sourceValue)
	case reflect.Array:
		return FromSlice(sourceValue)
	default:
		panic("seq.New can only be called on Seq, slices, or chans.")
	}
}
