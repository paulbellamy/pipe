package pipe

import (
	"reflect"
)

// ForEachSlice is of type: func(fn func(T), input []T).
// Execute a function for each item. Returns nothing.
func ForEachSlice(fn, input interface{}) {
	checkForEachFuncType(fn, input)

	fnValue := reflect.ValueOf(fn)
	inputValue := reflect.ValueOf(input)

	for i := 0; i < inputValue.Len(); i++ {
		fnValue.Call([]reflect.Value{inputValue.Index(i)})
	}
}
