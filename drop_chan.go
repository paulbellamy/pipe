package pipe

import (
  "fmt"
  "reflect"
)

// DropChan is of type: func(num int, input chan T) chan T.
// Drop a given number of items from the input chan. After that number has been
// dropped, the rest are passed straight through.
func DropChan(num int, input interface{}) interface{} {
  inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Chan {
		panic(fmt.Sprintf("DropChan called on invalid type: %s", inputValue.Type()))
	}

	output := reflect.MakeChan(inputValue.Type(), 0)
	var count int
	go func() {
		// drop num items
		for count = 0; count < num; count++ {
			_, ok := inputValue.Recv()
			if !ok {
				// channel closed early
				output.Close()
				return
			}
		}

		// Return the rest
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
