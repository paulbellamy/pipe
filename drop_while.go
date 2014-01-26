package pipe

import (
	"fmt"
	"reflect"
)

func checkDropWhileFuncType(fn, input interface{}) {
	fnType := reflect.TypeOf(fn)
	inputType := reflect.TypeOf(input)

	valid := fnType.NumOut() == 1 &&
		fnType.NumIn() == 1 &&
		inputType.Elem().ConvertibleTo(fnType.In(0)) &&
		fnType.Out(0).ConvertibleTo(boolType)

	if !valid {
		panic(fmt.Sprintf("DropWhile fn must be of type func(%v) bool, but was %v", inputType.Elem(), fnType))
	}
}
