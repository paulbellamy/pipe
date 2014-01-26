package pipe

import (
	"reflect"
)

func FieldByName(name string) func(interface{}) interface{} {
	return func(record interface{}) interface{} {
		return reflect.ValueOf(record).FieldByName(name).Interface()
	}
}
