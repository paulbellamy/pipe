package pipe

import (
	"fmt"
	"reflect"
)

// DropWhileChan is of type: func(fn func(T) bool, input chan T) chan T.
// Drop the items from the input chan until the given function returns true.
// After that, the rest are passed straight through.
func DropWhileChan(fn, input interface{}) interface{} {
	checkDropWhileFuncType(fn, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	if inputValue.Kind() != reflect.Chan {
		panic(fmt.Sprintf("DropWhileChan called on invalid type: %s", inputValue.Type()))
	}

	output := reflect.MakeChan(inputValue.Type(), 0)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				// input closed, abort
				output.Close()
				return
			}

			// check if we should output this
			if !fnValue.Call([]reflect.Value{item})[0].Bool() {
				output.Send(item)
				break
			}
		}

		// send any messages after this
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			output.Send(item)
		}

		output.Close()

	}()
	return output.Interface()
}
