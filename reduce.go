// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "reflect"
)

// Accumulate the result of the reduce function being called on each item, then
// when the input channel is closed, pass the result to the output channel
func Reduce(input interface{}, initial interface{}, fn interface{}) interface{} {
  initialType := reflect.TypeOf(initial)
	inputValue := reflect.ValueOf(input)
  inputType := inputValue.Type()
	fnValue := reflect.ValueOf(fn)

  signature := &functionSignature{
    []reflect.Type{initialType, inputType.Elem()},
    []reflect.Type{initialType},
  }
  signature.Check("Reduce fn", fn)

  outputType := reflect.ChanOf(reflect.BothDir, initialType)
	output := reflect.MakeChan(outputType, 0)

  result := reflect.ValueOf(initial)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			result = fnValue.Call([]reflect.Value{result, item})[0]
		}
		// Input was closed, send the result
		output.Send(result)
    output.Close()
	}()
	return output.Interface()
}
