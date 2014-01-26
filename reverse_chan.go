package pipe

import (
	"fmt"
	"reflect"
)

// ReverseChan is of type: func(input chan T) chan T.
// Reverse the order of items flowing on a chan. Builds up a list in memory of
// all the values until the input channel is closed, then emits all received
// values. Can be memory intensive for large chans.
func ReverseChan(input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Chan {
		panic(fmt.Sprintf("ReverseChan called on invalid type: %s", inputValue.Type()))
	}

	output := reflect.MakeChan(inputValue.Type(), 0)
	stored := reflect.MakeSlice(reflect.SliceOf(inputValue.Type().Elem()), 0, 1)
	go func() {
		count := 0
		// Build a list of received items
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			stored = reflect.Append(stored, item)
			count++
		}

		// Send them all out in reverse order
		for count--; count >= 0; count-- {
			output.Send(stored.Index(count))
		}
		output.Close()
	}()
	return output.Interface()
}
