// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
  "fmt"
  "reflect"
)

// Group each message from the input channel with it's corresponding message
// from the other channel. This will block on the first channel until it
// receives a message, then block on the second until it gets one from there.
// At that point an array containing both will be sent to the output channel.
//
// For example, if channel a is being zipped with channel b, and output on channel c:
//
//   a <- 1
//   b <- 2
//   result := <-c // result will equal []int{1, 2}
//

func zipValues(t reflect.Type, a, b reflect.Value) reflect.Value {
  zipped := reflect.MakeSlice(t, 0, 2)
  return reflect.Append(zipped, a, b)
}

func Zip(input interface{}, other interface{}) interface{} {
	inputValue := reflect.ValueOf(input)
  inputType := inputValue.Type()

  otherValue := reflect.ValueOf(other)
  otherType := otherValue.Type()

  if inputType != otherType {
    panic(fmt.Sprintf("Zip input types must match, but they were %v and %v", inputType, otherType))
  }

  zippedType := reflect.SliceOf(inputType.Elem())
  outputType := reflect.ChanOf(reflect.BothDir, zippedType)
	output := reflect.MakeChan(outputType, 0)
	go func() {
		// only send num items
		for {
			a, ok := inputValue.Recv()
			if !ok {
				break
			}

			b, ok := otherValue.Recv()
			if !ok {
				break
			}

      output.Send(zipValues(zippedType, a, b))
		}

    output.Close()
	}()
	return output.Interface()
}
