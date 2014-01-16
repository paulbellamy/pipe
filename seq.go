package pipe

import (
	"reflect"
)

type Seq interface {
	First() interface{}
	Rest() Seq
	Empty() bool
}

// Create a new Seq. Supports Chans and Slices.
func New(source interface{}) Seq {
	sourceValue := ensureValue(source)
	switch sourceValue.Type().Kind() {
	case reflect.Chan:
		return FromChan(sourceValue)
	default:
		return FromSlice(sourceValue)
	}
}

// Make sure that x is a reflect.Value. Convert it to one if it isn't.
func ensureValue(x interface{}) reflect.Value {
	if reflect.TypeOf(x).String() == "reflect.Value" {
		return x.(reflect.Value)
	} else {
		return reflect.ValueOf(x)
	}
}
