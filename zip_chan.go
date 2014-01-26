package pipe

import (
	"fmt"
	"reflect"
)

// ZipChan is of type: func(input chan T, others ...chan T) chan []T.
// Group each message from the input channel with it's corresponding message(s)
// from the other channel(s). This will block on the first channel until it
// receives a message, then block on the second until it gets one from there.
// At that point an array containing both will be sent to the output channel.
func ZipChan(input interface{}, others ...interface{}) interface{} {
	inputType := reflect.TypeOf(input)

	inputs := append([]interface{}{input}, others...)
	inputValues := MapSlice(reflect.ValueOf, inputs).([]reflect.Value)

	for i := 0; i < len(inputValues); i++ {
		if inputValues[i].Kind() != reflect.Chan {
			panic(fmt.Sprintf("ZipChan called on invalid type: %s", inputValues[i].Type()))
		}

		if inputValues[i].Type() != inputType {
			panic(fmt.Sprintf("Zip input types must match, but they were %v and %v", inputType, inputValues[i].Type()))
		}
	}

	zippedType := reflect.SliceOf(inputType.Elem())
	outputType := reflect.ChanOf(reflect.BothDir, zippedType)
	output := reflect.MakeChan(outputType, 0)
	go func() {
		i := 0
		for {
			zipped := reflect.MakeSlice(zippedType, 0, len(inputValues))

			for i = 0; i < len(inputValues); i++ {
				item, ok := inputValues[i].Recv()
				if !ok {
					break
				}
				zipped = reflect.Append(zipped, item)
			}

			if i < len(inputValues) {
				break
			}

			output.Send(zipped)
		}

		output.Close()
	}()
	return output.Interface()
}
