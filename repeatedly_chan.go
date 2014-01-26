package pipe

import (
	"reflect"
)

func RepeatedlyChan(fn interface{}) interface{} {
	fnValue := reflect.ValueOf(fn)

	outputType := reflect.ChanOf(reflect.BothDir, fnValue.Type().Out(0))
	output := reflect.MakeChan(outputType, 0)

	go func() {
		args := []reflect.Value{}
		for {
			output.Send(fnValue.Call(args)[0])
		}
	}()
	return output.Interface()
}
