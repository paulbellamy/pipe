// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"reflect"
)

// Execute a function for each item (without modifying the item). Useful for
// monitoring, logging, or causing some side-effect.
func ForEach(fn, input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)
	inputType := inputValue.Type()
	fnValue := reflect.ValueOf(fn)

	signature := &functionSignature{
		[]reflect.Type{inputType.Elem()},
		[]reflect.Type{},
	}
	signature.Check("ForEach fn", fn)

	output := reflect.MakeChan(inputType, 0)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			fnValue.Call([]reflect.Value{item})
			output.Send(item)
		}
		output.Close()
	}()
	return output.Interface()
}
