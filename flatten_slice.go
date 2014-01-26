package pipe

import (
	"fmt"
	"reflect"
)

// FlattenSlice is of type: func(input [][]T) []T.
// Takes a chan of arrays, and concatenates them together, putting each element
// onto the output chan. After input is closed, output is also closed. If input
// is []T instead of type [][]T, then this is a no-op.
func FlattenSlice(input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Slice &&
		inputValue.Kind() != reflect.Array {
		panic(fmt.Sprintf("FlattenSlice called on invalid type: %s", inputValue.Type()))
	}

	elemType := inputValue.Type().Elem()
	if elemType.Kind() != reflect.Array &&
		elemType.Kind() != reflect.Slice {
		return input
	}

	outputType := reflect.SliceOf(elemType.Elem())
	output := reflect.MakeSlice(outputType, 0, 1)

	for i := 0; i < inputValue.Len(); i++ {
		items := inputValue.Index(i)
		for j := 0; j < items.Len(); j++ {
			output = reflect.Append(output, items.Index(j))
		}
	}
	return output.Interface()
}
