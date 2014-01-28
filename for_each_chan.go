package pipe

import (
	"reflect"
)

// ForEachChan is of type: func(fn func(T), input chan T)
// Execute a function for each item. Useful
// for monitoring, logging, or causing some side-effect. Returns nothing
func ForEachChan(fn, input interface{}) {
	checkForEachFuncType(fn, input)

	fnValue := reflect.ValueOf(fn)
	inputValue := reflect.ValueOf(input)

	for {
		item, ok := inputValue.Recv()
		if !ok {
			break
		}

		fnValue.Call([]reflect.Value{item})
	}
}
