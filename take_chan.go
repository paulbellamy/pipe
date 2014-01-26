package pipe

import (
	"fmt"
	"reflect"
)

// TakeChan is of type: func(num int, input chan T) chan T.
// Accept only the given number of items from the input chan. After that number
// has been received, all input messages will be ignored and the output channel
// will be closed.
func TakeChan(num int, input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Chan {
		panic(fmt.Sprintf("DropChan called on invalid type: %s", inputValue.Type()))
	}

	output := reflect.MakeChan(inputValue.Type(), 0)
	var count int
	go func() {
		// only send num items
		for count = 0; count < num; count++ {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			output.Send(item)
		}

		// sent our max, close the channel
		output.Close()
	}()
	return output.Interface()
}
