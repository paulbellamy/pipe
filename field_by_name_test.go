package pipe

import (
	"testing"
)

type customer struct {
	Name string
}

func TestFieldByName(t *testing.T) {
	alice := customer{"Alice"}

	getter := FieldByName("Name")
	result := getter(alice).(string)
	if result != "Alice" {
		t.Fatal("Expected getter to return 'Alice', but returned,", result)
	}
}
