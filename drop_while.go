// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "reflect"
)

// Drop the items from the input pipe until the given function returns true.
// After that , the rest are passed straight through.
func DropWhile(input interface{}, fn interface{}) interface{} {
	inputValue := reflect.ValueOf(input)
  inputType := inputValue.Type()
	fnValue := reflect.ValueOf(fn)

  signature := &functionSignature{
    []reflect.Type{inputType.Elem()},
    []reflect.Type{reflect.TypeOf(false)},
  }
  signature.Check("DropWhile fn", fn)

	output := reflect.MakeChan(inputType, 0)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				// input closed, abort
        output.Close()
				return
			}

			// check if we should output this
			if !fnValue.Call([]reflect.Value{item})[0].Bool() {
        output.Send(item)
				break
			}
		}

		// send any messages after this
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

      output.Send(item)
		}

    output.Close()

	}()
	return output.Interface()
}
