// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
	"reflect"
)

func checkMapFuncType(fn, input interface{}) {
	inputType := reflect.TypeOf(input)
	fnType := reflect.TypeOf(fn)

	valid := fnType.NumOut() == 1 &&
		fnType.NumIn() == 1 &&
		inputType.Elem().ConvertibleTo(fnType.In(0))

	if !valid {
		panic(fmt.Sprintf("Map fn must be of type func(%v) T, but was %v", inputType.Elem(), fnType))
	}
}
