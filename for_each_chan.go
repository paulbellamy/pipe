package pipe

import (
	"reflect"
)

// ForEachChan is of type: func(fn func(T), input chan T) chan T.
// Execute a function for each item (without modifying the item). Useful
// for monitoring, logging, or causing some side-effect. Returns a channel receiving the input.
func ForEachChan(fn, input interface{}) interface{} {
  checkForEachFuncType(fn, input)

	fnValue := reflect.ValueOf(fn)
	inputValue := reflect.ValueOf(input)

	output := reflect.MakeChan(inputValue.Type(), 0)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			fnValue.Call([]reflect.Value{item})
			output.Send(item)
		}
		output.Close()
	}()
	return output.Interface()
}
