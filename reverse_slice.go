package pipe

import (
	"fmt"
	"reflect"
)

// ReverseSlice is of type: func(input []T) []T.
// Reverse the order of items in a slice. Reverses the items in place,
// modifying the original array. The array is returned.
func ReverseSlice(input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Slice &&
		inputValue.Kind() != reflect.Array {
		panic(fmt.Sprintf("ReverseSlice called on invalid type: %s", inputValue.Type()))
	}

	var temp reflect.Value
	length := inputValue.Len()
	a := 0
	b := length - 1
	for a < b {
		fmt.Println(inputValue.Index(a).Int(), inputValue.Index(b).Int())
		temp = inputValue.Index(a)
		inputValue.Index(a).Set(inputValue.Index(b))
		inputValue.Index(b).Set(temp)
		fmt.Println(inputValue.Index(a).Int(), inputValue.Index(b).Int())
		a++
		b--
	}

	return input
}
