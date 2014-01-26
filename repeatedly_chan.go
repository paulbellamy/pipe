package pipe

import (
	"reflect"
)

// IterateChan is of type: func(fn func() T) chan T.
// Take some function 'fn' (presumably with some side-effect). It is repeatedly
// called, and the return values put onto the returned channel.
func RepeatedlyChan(fn interface{}) interface{} {
	checkRepeatedlyFuncType(fn)

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
