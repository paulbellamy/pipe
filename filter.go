// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
	"reflect"
)

var boolType = reflect.TypeOf(true)

func checkFilterFuncType(fn, input interface{}) {
	fnType := reflect.TypeOf(fn)
	inputType := reflect.TypeOf(input)

	valid := fnType.NumOut() == 1 &&
		fnType.NumIn() == 1 &&
		inputType.Elem().ConvertibleTo(fnType.In(0)) &&
		fnType.Out(0).ConvertibleTo(boolType)

	if !valid {
		panic(fmt.Sprintf("Filter fn must be of type func(%v) bool, but was %v", inputType.Elem(), fnType))
	}
}
