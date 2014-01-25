// Copyright 2014 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
	"reflect"
)

func checkForEachFuncType(fn, input interface{}) {
	fnType := reflect.TypeOf(fn)
	inputType := reflect.TypeOf(input)

	valid := fnType.NumIn() == 1
  if fnType.IsVariadic() {
    valid = valid && inputType.Elem().ConvertibleTo(fnType.In(0).Elem())
  } else {
    valid = valid && inputType.Elem().ConvertibleTo(fnType.In(0))
  }

	if !valid {
		panic(fmt.Sprintf("ForEach fn must be of type func(%v), but was %v", inputType.Elem(), fnType))
	}
}
