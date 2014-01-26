package pipe

import (
	"reflect"
)

// ReduceChan is of type: func(fn func(accumulator U, item T) U, initial U, input chan T) U.
// It accumulates the result of the reduce function being called on each item,
// then when the input channel is closed, return the result.
func ReduceChan(fn, initial, input interface{}) interface{} {
	checkReduceFuncType(fn, initial, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	result := reflect.ValueOf(initial)
	for {
		item, ok := inputValue.Recv()
		if !ok {
			break
		}

		result = fnValue.Call([]reflect.Value{result, item})[0]
	}
	return result.Interface()
}
