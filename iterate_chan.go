package pipe

import (
	"reflect"
)

func IterateChan(fn interface{}, initialArgs ...interface{}) interface{} {
	fnValue := reflect.ValueOf(fn)
	initialValues := MapSlice(reflect.ValueOf, initialArgs).([]reflect.Value)

	checkIterateFuncType(fnValue, initialValues)

	outputType := reflect.ChanOf(reflect.BothDir, fnValue.Type().Out(0))
	output := reflect.MakeChan(outputType, 0)

	go func() {
		args := initialValues
		for {
			output.Send(args[0])
			args = fnValue.Call(args)
		}
	}()
	return output.Interface()
}
