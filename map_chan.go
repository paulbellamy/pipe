package pipe

import (
	"reflect"
)

// MapChan is of type: func(fn func(T) U, input chan T) chan U.
// It returns a chan which receives fn(item) for each item in input.
func MapChan(fn, input interface{}) interface{} {
	checkMapFuncType(fn, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	outputType := reflect.ChanOf(reflect.BothDir, fnValue.Type().Out(0))
	output := reflect.MakeChan(outputType, 0)
	go func() {
		for {
			value, ok := inputValue.Recv()
			if !ok {
				break
			}

			output.Send(fnValue.Call([]reflect.Value{value})[0])
		}
		output.Close()
	}()
	return output.Interface()
}
