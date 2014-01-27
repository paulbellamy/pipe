package pipe

import (
	"reflect"
)

func interfaceOf(val reflect.Value) interface{} {
	return val.Interface()
}

func Call(args ...interface{}) func(interface{}) []interface{} {
	argValues := MapSlice(reflect.ValueOf, args).([]reflect.Value)
	return func(fn interface{}) []interface{} {
		return MapSlice(interfaceOf, reflect.ValueOf(fn).Call(argValues)).([]interface{})
	}
}
