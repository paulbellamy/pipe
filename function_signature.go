package pipe

import (
	"fmt"
	"reflect"
	"strings"
)

type functionSignature struct {
	inputs  []reflect.Type
	outputs []reflect.Type
}

func (f *functionSignature) Check(name string, actual interface{}) {
  actualType := reflect.TypeOf(actual)
	if !f.ConvertibleTo(actualType) {
		panic(fmt.Sprintf("%v must be of type %v, but was %v", name, f, actualType))
	}
}


func (f *functionSignature) Align() int {
  return 8
}

func (f *functionSignature) FieldAlign() int {
  return 8
}

func (f *functionSignature) Method(int) reflect.Method {
  panic(fmt.Sprintf("%s has no methods", f))
}

func (f *functionSignature) MethodByName(string) (reflect.Method, bool) {
  return reflect.Method{}, false
}

func (f *functionSignature) NumMethod() int {
  return 0
}

func (f *functionSignature) Name() string {
  return ""
}

func (f *functionSignature) PkgPath() string {
  return ""
}

func (f *functionSignature) Size() uintptr {
  return 8
}

func formatTypesList(types []reflect.Type) string {
  output := []string{}
  for i := 0; i < len(types); i++ {
    output = append(output, types[i].String())
  }
  return strings.Join(output, ", ")
}

func (f *functionSignature) String() string {
  inputsList := formatTypesList(f.inputs)

  outputsList := ""
  if len(f.outputs) > 1 {
    outputsList = fmt.Sprintf(" (%v)", formatTypesList(f.outputs))
  } else if len(f.outputs) > 0 {
    outputsList = fmt.Sprintf(" %v", formatTypesList(f.outputs))
  }

  return fmt.Sprintf("func(%s)%s", inputsList, outputsList)
}

func (f *functionSignature) Kind() reflect.Kind {
  return reflect.Func
}

func (f *functionSignature) Implements(reflect.Type) bool {
  return false
}

func (f *functionSignature) AssignableTo(other reflect.Type) bool {
  return f.ConvertibleTo(other)
}

func (f *functionSignature) ConvertibleTo(other reflect.Type) bool {
  if other.IsVariadic() {
    // No support for variadic functions yet.
    return false
  }

  if other.NumIn() != len(f.inputs) {
    return false
  }
  for i := 0; i < len(f.inputs); i++ {
    if f.inputs[i] != other.In(i) {
      return false
    }
  }

  if other.NumOut() != len(f.outputs) {
    return false
  }
  for i := 0; i < len(f.outputs); i++ {
    if f.outputs[i] != other.Out(i) {
      return false
    }
  }

  return true
}

func (f *functionSignature) Bits() int {
  panic(fmt.Sprintf("Bits of non-arithmetic Type %v", f))
}

func (f *functionSignature) ChanDir() int {
  panic("ChanDir of non-chan type")
}

func (f *functionSignature) IsVariadic() bool {
  return false
}

func (f *functionSignature) Elem() reflect.Type {
  panic("Elem of invalid type")
}

func (f *functionSignature) Field(int) reflect.StructField {
  panic("Field of non-struct type")
}

func (f *functionSignature) FieldByName(string) (reflect.StructField, bool) {
  panic("FieldByName of non-struct type")
}

func (f *functionSignature) FieldByNameFunc(func(string) bool) (reflect.StructField, bool) {
  panic("FieldByNameFunc of non-struct type")
}

func (f *functionSignature) In(i int) reflect.Type {
  return f.inputs[i]
}

func (f *functionSignature) Key() reflect.Type {
  panic("Key of non-map type")
}

func (f *functionSignature) Len() reflect.Type {
  panic("Key of non-array type")
}

func (f *functionSignature) NumField() int {
  panic("NumField of non-struct type")
}

func (f *functionSignature) NumIn() int {
  return len(f.inputs)
}

func (f *functionSignature) NumOut() int {
  return len(f.outputs)
}

func (f *functionSignature) Out(i int) reflect.Type {
  return f.outputs[i]
}
