package pipe

import (
	"fmt"
	"reflect"
)

// DropWhileSlice is of type: func(fn func(T) bool, input []T) []T.
// Drop the items from the input slice until the given function returns true.
// After that, the rest are passed straight through.
func DropWhileSlice(fn, input interface{}) interface{} {
	checkFilterFuncType(fn, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	if inputValue.Kind() != reflect.Slice &&
		inputValue.Kind() != reflect.Array {
		panic(fmt.Sprintf("DropWhileChan called on invalid type: %s", inputValue.Type()))
	}

	i := 0
	for ; i < inputValue.Len(); i++ {
		if !fnValue.Call([]reflect.Value{inputValue.Index(i)})[0].Bool() {
			break
		}
	}

	return inputValue.Slice(i, inputValue.Len()).Interface()
}
