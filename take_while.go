// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"reflect"
)

// Accept items from the input pipe until the given function returns false.
// After that, all input messages will be ignored and the output channel will
// be closed.
func TakeWhile(fn, input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)
	inputType := inputValue.Type()
	fnValue := reflect.ValueOf(fn)

	signature := &functionSignature{
		[]reflect.Type{inputType.Elem()},
		[]reflect.Type{reflect.TypeOf(false)},
	}
	signature.Check("TakeWhile fn", fn)

	output := reflect.MakeChan(inputType, 0)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			// check if we should continue
			if !fnValue.Call([]reflect.Value{item})[0].Bool() {
				break
			}

			output.Send(item)
		}

		// hit the toggle, close the channel
		output.Close()

		// drop any extra messages
		for {
			_, ok := inputValue.Recv()
			if !ok {
				break
			}
		}
	}()
	return output.Interface()
}
