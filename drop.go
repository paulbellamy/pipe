// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"reflect"
)

// Drop a given number of items from the input pipe. After that number has been
// dropped, the rest are passed straight through.
func Drop(input interface{}, num int) interface{} {
	inputValue := reflect.ValueOf(input)

	switch inputValue.Kind() {
	case reflect.Chan:
		return dropChan(inputValue, num)
	case reflect.Array:
		return dropSlice(inputValue, num)
	case reflect.Slice:
		return dropSlice(inputValue, num)
	}
	panic("Drop called on invalid type")
}

func dropChan(input reflect.Value, num int) interface{} {
	inputType := input.Type()

	output := reflect.MakeChan(inputType, 0)
	var count int
	go func() {
		// drop num items
		for count = 0; count < num; count++ {
			_, ok := input.Recv()
			if !ok {
				// channel closed early
				output.Close()
				return
			}
		}

		// Return the rest
		for {
			item, ok := input.Recv()
			if !ok {
				break
			}

			output.Send(item)
		}
		output.Close()
	}()
	return output.Interface()
}

func dropSlice(input reflect.Value, num int) interface{} {
	if num > input.Len() {
		return input.Slice(0, 0).Interface()
	} else {
		return input.Slice(num, input.Len()).Interface()
	}
}
