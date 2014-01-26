package pipe

import (
	"reflect"
)

// TakeWhileChan is of type: func(fn func(T) bool, input chan T) chan T.
// Accept items from the input chan until the given function returns false.
// After that, all input messages will be ignored and the output channel will
// be closed.
func TakeWhileChan(fn, input interface{}) interface{} {
	checkTakeWhileFuncType(fn, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	output := reflect.MakeChan(inputValue.Type(), 0)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			// check if we should continue
			if !fnValue.Call([]reflect.Value{item})[0].Bool() {
				break
			}

			output.Send(item)
		}

		// hit the toggle, close the channel
		output.Close()

		// drop any extra messages
		for {
			_, ok := inputValue.Recv()
			if !ok {
				break
			}
		}
	}()
	return output.Interface()
}
