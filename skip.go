// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "reflect"
)

// Skip a given number of items from the input pipe. After that number has been
// dropped, the rest are passed straight through.
func Skip(input interface{}, num int64) interface{} {
	inputValue := reflect.ValueOf(input)
  inputType := inputValue.Type()

	output := reflect.MakeChan(inputType, 0)
	var count int64
	go func() {
		// skip num items
		for count = 0; count < num; count++ {
			_, ok := inputValue.Recv()
			if !ok {
				// channel closed early
        output.Close()
				return
			}
		}

		// Return the rest
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
