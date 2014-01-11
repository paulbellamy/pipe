package pipe

import (
	"reflect"
	"testing"
)

var intType = reflect.TypeOf(5)
var boolType = reflect.TypeOf(true)
var sig = &functionSignature{[]reflect.Type{intType}, []reflect.Type{boolType}}

func TestFunctionSignature(t *testing.T) {
	sig.Check("MyFunc", func(int) bool { return true })

	invalids := []interface{}{
		func(int, int) bool { return true },
		func() bool { return true },
		func(int) (bool, int) { return true, 1 },
		func(int) {},
	}
	test_invalid := func(invalid interface{}) {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected to panic on invalid type: %v", reflect.TypeOf(invalid))
			}
		}()

		sig.Check("MyFunc", invalid)
	}
	for i := 0; i < len(invalids); i++ {
		test_invalid(invalids[i])
	}
}
