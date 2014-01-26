package pipe

import (
	"fmt"
	"reflect"
)

// FlattenChan is of type: func(input chan []T) chan T.
// Takes a chan of arrays, and concatenates them together, putting each element
// onto the output chan. After input is closed, output is also closed. If input
// is chan T instead of type chan []T, then this is a no-op.
func FlattenChan(input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Chan {
		panic(fmt.Sprintf("FlattenChan called on invalid type: %s", inputValue.Type()))
	}

	elemType := inputValue.Type().Elem()
	if elemType.Kind() != reflect.Array &&
		elemType.Kind() != reflect.Slice {
		return input
	}

	outputType := reflect.ChanOf(reflect.BothDir, elemType.Elem())
	output := reflect.MakeChan(outputType, 0)
	go func() {
		for {
			value, ok := inputValue.Recv()
			if !ok {
				break
			}

			for i := 0; i < value.Len(); i++ {
				output.Send(value.Index(i))
			}
		}
		output.Close()
	}()
	return output.Interface()
}
