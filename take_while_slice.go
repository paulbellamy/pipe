package pipe

import (
	"fmt"
	"reflect"
)

// TakeWhileSlice is of type: func(fn func(T) bool, input []T) []T.
// Accept items from the input slice until the given function returns false.
// After that, the rest are passed straight through.
func TakeWhileSlice(fn, input interface{}) interface{} {
	checkTakeWhileFuncType(fn, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	if inputValue.Kind() != reflect.Slice &&
		inputValue.Kind() != reflect.Array {
		panic(fmt.Sprintf("TakeWhileChan called on invalid type: %s", inputValue.Type()))
	}

	i := 0
	for ; i < inputValue.Len(); i++ {
		if !fnValue.Call([]reflect.Value{inputValue.Index(i)})[0].Bool() {
			break
		}
	}

	return inputValue.Slice(0, i).Interface()
}
