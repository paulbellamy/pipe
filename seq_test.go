package pipe

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, a, b interface{}) {
	if reflect.ValueOf(a) != reflect.ValueOf(b) {
		t.Errorf("Expected: %v (Type %v)\nActual: %v (Type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

func assert(t *testing.T, x bool) {
	assertEqual(t, true, x)
}

func TestSeqAcceptsSlice(t *testing.T) {
	source := []int{1, 2, 3}
	s := New(source)
	assertEqual(t, 1, s.First())
}

func TestSeqAcceptsChan(t *testing.T) {
	source := make(chan int)
	s := New(source)
	go func() {
		source <- 1
	}()
	assertEqual(t, 1, s.First())
}

func TestSeqAcceptsReflectValue(t *testing.T) {
	source := []int{1, 2, 3}
	s := New(reflect.ValueOf(source))
	assertEqual(t, 1, s.First())
}

func TestSeqPanicsOnUnknownType(t *testing.T) {
  t.Fatal("Not tested")
}
