package pipe

import (
  "fmt"
  "reflect"
)

func MapSlice(fn, input interface{}) interface{} {
	checkMapFuncType(fn, input)

  inputValue := reflect.ValueOf(input)
  fnValue := reflect.ValueOf(fn)

	if inputValue.Kind() != reflect.Slice &&
		inputValue.Kind() != reflect.Array {
      panic(fmt.Sprintf("MapSlice called on invalid type: %s", inputValue.Type()))
	}

	outputType := reflect.SliceOf(fnValue.Type().Out(0))
	output := reflect.MakeSlice(outputType, 0, inputValue.Len())

	for i := 0; i < inputValue.Len(); i++ {
		output = reflect.Append(
			output,
			fnValue.Call([]reflect.Value{inputValue.Index(i)})[0],
		)
	}

	return output.Interface()
}
