// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
	"reflect"
	"seq"
)

func Map(fn, input interface{}) interface{} {
	checkMapFuncType(fn, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	switch inputValue.Kind() {
	case reflect.Chan:
		return mapChan(fnValue, inputValue)
	case reflect.Array:
		return mapSlice(fnValue, inputValue)
	case reflect.Slice:
		return mapSlice(fnValue, inputValue)
	case reflect.Map:
		return mapMap(fnValue, inputValue)
	}
	panic("Map called on invalid type")
}

// Pass through the result of the map function for each item
func mapChan(fn, input reflect.Value) interface{} {
	outputType := reflect.ChanOf(reflect.BothDir, fn.Type().Out(0))
	output := reflect.MakeChan(outputType, 0)
	go func() {
		for s := seq.New(input); !s.Empty(); s = s.Rest() {
			output.Send(fn.Call([]reflect.Value{reflect.ValueOf(s.First())})[0])
		}
		output.Close()
	}()
	return output.Interface()
}

func mapSlice(fn, input reflect.Value) interface{} {
	outputType := reflect.SliceOf(fn.Type().Out(0))
	output := reflect.MakeSlice(outputType, 0, input.Len())

	for s := seq.New(input); !s.Empty(); s = s.Rest() {
		output = reflect.Append(
			output,
			fn.Call([]reflect.Value{reflect.ValueOf(s.First())})[0],
		)
	}

	return output.Interface()
}

func mapMap(fn, input reflect.Value) interface{} {
	outputType := reflect.SliceOf(fn.Type().Out(0))
	output := reflect.MakeSlice(outputType, 0, input.Len())
	for _, key := range input.MapKeys() {
		output = reflect.Append(output, fn.Call([]reflect.Value{key, input.MapIndex(key)})[0])
	}

	return output.Interface()
}

func checkMapFuncType(fn, input interface{}) {
	inputType := reflect.TypeOf(input)
	fnType := reflect.TypeOf(fn)

	valid := (fnType.NumOut() == 1)
	switch inputType.Kind() {
	case reflect.Map:
		valid = fnType.NumIn() == 2 &&
			fnType.In(0) == inputType.Key() &&
			fnType.In(1) == inputType.Elem()
	default:
		valid = fnType.NumIn() == 1 &&
			fnType.In(0) == inputType.Elem()
	}

	if !valid {
		panic(fmt.Sprintf("Map fn must be of type func(%v) T, but was %v", inputType.Elem(), fnType))
	}
}
