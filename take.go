// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "reflect"
)

// Accept only the given number of items from the input pipe. After that number
// has been received, all input messages will be ignored and the output channel
// will be closed.
func Take(input interface{}, num int64) interface{} {
	inputValue := reflect.ValueOf(input)
  inputType := inputValue.Type()

	output := reflect.MakeChan(inputType, 0)
	var count int64
	go func() {
		// only send num items
		for count = 0; count < num; count++ {
			item, ok := inputValue.Recv()
			if !ok {
				break
			}

      output.Send(item)
		}

		// sent our max, close the channel
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
