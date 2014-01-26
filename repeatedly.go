package pipe

import (
	"fmt"
	"reflect"
)

func checkRepeatedlyFuncType(fn interface{}) {
	fnType := reflect.TypeOf(fn)

	valid := fnType.NumOut() == 1 && fnType.NumIn() == 0

	if !valid {
		panic(fmt.Sprintf("Repeatedly fn must be of type func() T, but was %v", fnType))
	}
}
