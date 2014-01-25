package pipe

import (
	"reflect"
)

// ForEachSlice is of type: func(fn func(T), input []T) []T.
// Execute a function for each item (without modifying the item). Useful
// for monitoring, logging, or causing some side-effect. Returns the
// original input.
func ForEachSlice(fn, input interface{}) interface{} {
  checkForEachFuncType(fn, input)

	fnValue := reflect.ValueOf(fn)
	inputValue := reflect.ValueOf(input)

  for i := 0; i < inputValue.Len(); i++ {
    fnValue.Call([]reflect.Value{inputValue.Index(i)})
  }

  return input
}
