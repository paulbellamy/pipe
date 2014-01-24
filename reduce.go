// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
	"reflect"
)

func checkReduceFuncType(fn, initial, input interface{}) {
	fnType := reflect.TypeOf(fn)
	initialType := reflect.TypeOf(initial)
	inputType := reflect.TypeOf(input)

	valid := fnType.NumOut() == 1 &&
		fnType.NumIn() == 2 &&
		initialType.ConvertibleTo(fnType.In(0)) &&
		inputType.Elem().ConvertibleTo(fnType.In(1)) &&
		fnType.Out(0).ConvertibleTo(fnType.In(0))

	if !valid {
		panic(fmt.Sprintf("Reduce fn must be of type func(%v, %v) T, but was %v", initialType, inputType.Elem(), fnType))
	}
}
