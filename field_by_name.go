package pipe

import (
	"reflect"
)

// FieldByName is a helper for fetching a named field from a struct. It is
// useful in conjunction with Map, to retrieve many values at once.
func FieldByName(name string) func(interface{}) interface{} {
	return func(record interface{}) interface{} {
		return reflect.ValueOf(record).FieldByName(name).Interface()
	}
}
