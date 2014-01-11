// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "reflect"
)

// Apply a filtering function to a channel, which will only pass through items
// when the filter func returns true.
func Filter(input interface{}, fn interface{}) interface{} {
	inputValue := reflect.ValueOf(input)
  inputType := inputValue.Type()
	fnValue := reflect.ValueOf(fn)

  signature := &functionSignature{
    []reflect.Type{inputType.Elem()},
    []reflect.Type{reflect.TypeOf(false)},
  }
  signature.Check("Filter fn", fn)

	output := reflect.MakeChan(inputType, 0)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			if fnValue.Call([]reflect.Value{item})[0].Bool() {
				output.Send(item)
			}
		}
    output.Close()
	}()
	return output.Interface()
}
