package pipe

import (
	"reflect"
)

// MethodByName is a helper for fetching a named method on a struct. It is
// useful in conjunction with Map, to retrieve many values at once.
func MethodByName(name string) interface{} {
	return func(record interface{}) interface{} {
		return reflect.ValueOf(record).MethodByName(name).Interface()
	}
}
