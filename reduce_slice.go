package pipe

import (
	"reflect"
)

// ReduceSlice is of type: func(fn func(accumulator U, item T) U, initial U, input []T) U.
// It accumulates the result of the fn function being called on each item, with
// the last value being returned.
func ReduceSlice(fn, initial, input interface{}) interface{} {
	checkReduceFuncType(fn, initial, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	result := reflect.ValueOf(initial)
	for i := 0; i < inputValue.Len(); i++ {
		result = fnValue.Call([]reflect.Value{result, inputValue.Index(i)})[0]
	}
	return result.Interface()
}
