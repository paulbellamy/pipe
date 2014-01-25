package pipe

import (
	"fmt"
	"reflect"
)

// DropSlice is of type: func(num int, input []T) []T.
// Drop a given number of items from the input slice. After that number has been
// dropped, the rest are passed straight through.
func DropSlice(num int, input interface{}) interface{} {
  inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Slice &&
    inputValue.Kind() != reflect.Array {
		panic(fmt.Sprintf("DropSlice called on invalid type: %s", inputValue.Type()))
	}

	if num > inputValue.Len() {
		return inputValue.Slice(0, 0).Interface()
	} else {
		return inputValue.Slice(num, inputValue.Len()).Interface()
	}
}
