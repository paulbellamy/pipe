package pipe

import (
	"fmt"
	"reflect"
)

// FilterChan is of type: func(fn func(T) bool, input chan T) chan T.
// Apply a filtering function to a chan, which will only pass through
// items when the filter func returns true.
func FilterChan(fn, input interface{}) interface{} {
	checkFilterFuncType(fn, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	if inputValue.Kind() != reflect.Chan {
		panic(fmt.Sprintf("FilterChan called on invalid type: %s", inputValue.Type()))
	}

	output := reflect.MakeChan(inputValue.Type(), 0)
  go func() {
    for {
      item, ok := inputValue.Recv()
      if !ok {
        break
      }

      if fnValue.Call([]reflect.Value{item})[0].Bool() {
        output.Send(item)
      }
    }
    output.Close()
  }()

  return output.Interface()
}
