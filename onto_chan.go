package pipe

import (
	"fmt"
	"reflect"
)

// OntoChan is of type: func(input []T) chan T.
// Puts the elements of the input collection onto a channel, which is then
// closed.
func OntoChan(input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Slice &&
		inputValue.Kind() != reflect.Array {
		panic(fmt.Sprintf("OntoChan called on invalid type: %s", inputValue.Type()))
	}

	outputType := reflect.ChanOf(reflect.BothDir, inputValue.Type().Elem())
	output := reflect.MakeChan(outputType, 0)
	go func() {
		for i := 0; i < inputValue.Len(); i++ {
			output.Send(inputValue.Index(i))
		}
		output.Close()
	}()
	return output.Interface()
}
