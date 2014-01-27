package pipe

import (
  "reflect"
)

// Make sure that x is a reflect.Value. Convert it to one if it isn't.
func ensureValue(x interface{}) reflect.Value {
       if reflect.TypeOf(x).String() == "reflect.Value" {
               return x.(reflect.Value)
       } else {
               return reflect.ValueOf(x)
       }
}
