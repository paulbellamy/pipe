package pipe

import (
	"fmt"
	"reflect"
)

func checkIterateFuncType(fn reflect.Value, initialValues []reflect.Value) {
	fnType := fn.Type()

	argsCount := len(initialValues)
	valid := argsCount > 0 &&
		fnType.NumOut() == argsCount &&
		fnType.NumIn() == argsCount

	for i := 0; i < argsCount; i++ {
		if !initialValues[i].Type().ConvertibleTo(fnType.In(i)) ||
			!fnType.Out(i).ConvertibleTo(initialValues[i].Type()) {
			fmt.Println(initialValues[i].Type(), "=>", fnType.In(i), ":", initialValues[i].Type().ConvertibleTo(fnType.In(i)))
			fmt.Println(fnType.Out(i), "=>", initialValues[i].Type(), ":", fnType.Out(i).ConvertibleTo(initialValues[i].Type()))
			valid = false
			break
		}
	}

	if !valid {
		panic(fmt.Sprintf("Iterate fn must be of type func(T) T, func(T, U) (T, U), etc., but was %v", fnType))
	}
}
