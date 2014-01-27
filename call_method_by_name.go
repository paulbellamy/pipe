package pipe

// CallMethodByName is a helper for fetching a named method on a struct. It is
// useful in conjunction with Map, to retrieve many values at once.
func CallMethodByName(name string, args ...interface{}) interface{} {
	call := Call(args)
	method := MethodByName(name).(func(interface{}) interface{})
	return func(record interface{}) []interface{} {
		return call(method(record))
	}
}
