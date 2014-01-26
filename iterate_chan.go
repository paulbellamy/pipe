package pipe

import (
	"reflect"
)

// IterateChan is of type: func(fn func(T) T, initialArgs ...T) chan T.
// Returns a channel with the values of x, fn(x), fn(fn(x)), etc... fn can take
// multiple arguments, and return multiple values, but the types and counts of
// arguments and return values must match. Only the first return value of each
// call to fn will be printed.
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
