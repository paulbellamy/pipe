// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "fmt"
  "reflect"
)

// Pass through the result of the map function for each item
func Map(input interface{}, fn interface{}) interface{} {
  inputValue := reflect.ValueOf(input)
  inputType := inputValue.Type()
  fnValue := reflect.ValueOf(fn)
  fnType := fnValue.Type()
  if fnType.NumIn() != 1 || fnType.In(0) != inputType.Elem() || fnType.NumOut() != 1 {
		panic(fmt.Sprintf("Map fn must be of type func(%v) T, but was %v", inputType.Elem(), fnType))
  }

  outputType := reflect.ChanOf(reflect.BothDir, fnType.Out(0))
	output := reflect.MakeChan(outputType, 0)
	go func() {
		for {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

			output.Send(fnValue.Call([]reflect.Value{item})[0])
		}
    output.Close()
	}()
	return output.Interface()
}
