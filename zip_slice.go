package pipe

import (
	"fmt"
	"reflect"
)

// ZipSlice is of type: func(input []T, others ...[]T) [][]T.
// Group each item from the input with it's corresponding item(s) from the others.
func ZipSlice(input interface{}, others ...interface{}) interface{} {
	inputType := reflect.TypeOf(input)

	inputs := append([]interface{}{input}, others...)
	inputValues := MapSlice(reflect.ValueOf, inputs).([]reflect.Value)

	for i := 0; i < len(inputValues); i++ {
		if inputValues[i].Kind() != reflect.Slice &&
			inputValues[i].Kind() != reflect.Array {
			panic(fmt.Sprintf("ZipSlice called on invalid type: %s", inputValues[i].Type()))
		}

		if inputValues[i].Type() != inputType {
			panic(fmt.Sprintf("Zip input types must match, but they were %v and %v", inputType, inputValues[i].Type()))
		}
	}

	outputLength := inputValues[0].Len()
	for i := 0; i < len(inputValues); i++ {
		length := inputValues[i].Len()
		if length < outputLength {
			outputLength = length
		}
	}

	elemType := inputType.Elem()
	zippedType := reflect.SliceOf(elemType)
	outputType := reflect.SliceOf(zippedType)
	output := reflect.MakeSlice(outputType, 0, outputLength)

	for i := 0; i < outputLength; i++ {
		zipped := reflect.MakeSlice(zippedType, 0, len(inputValues))

		for j := 0; j < len(inputValues); j++ {
			zipped = reflect.Append(zipped, inputValues[j].Index(i))
		}

		output = reflect.Append(output, zipped)
	}

	return output.Interface()
}
