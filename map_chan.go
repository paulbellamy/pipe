package pipe

import (
  "reflect"
)

// Pass through the result of the map function for each item
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
