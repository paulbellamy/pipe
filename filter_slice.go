package pipe

import (
	"fmt"
	"reflect"
)

// FilterSlice is of type: func(fn func(T) bool, input []T) []T.
// Apply a filtering function to a slice, which will only pass through items
// when the filter func returns true.
func FilterSlice(fn, input interface{}) interface{} {
	checkFilterFuncType(fn, input)

	inputValue := reflect.ValueOf(input)
	fnValue := reflect.ValueOf(fn)

	if inputValue.Kind() != reflect.Slice &&
		inputValue.Kind() != reflect.Array {
		panic(fmt.Sprintf("FilterSlice called on invalid type: %s", inputValue.Type()))
	}

	outputType := reflect.SliceOf(inputValue.Type().Elem())
	output := reflect.MakeSlice(outputType, 0, inputValue.Len())

	for i := 0; i < inputValue.Len(); i++ {
		if fnValue.Call([]reflect.Value{inputValue.Index(i)})[0].Bool() {
			output = reflect.Append(output, inputValue.Index(i))
		}
	}

	return output.Interface()
}
