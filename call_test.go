package pipe

import (
	"testing"
)

func TestCall(t *testing.T) {
	fn := func(a int, bs ...int) int {
		return ReduceSlice(func(a, b int) int { return a + b }, a, bs).(int)
	}

	result := Call(1, 2, 3)(fn)[0].(int)
	expected := 6
	if result != expected {
		t.Fatal("Expected call to return", expected, "but returned,", result)
	}
}
