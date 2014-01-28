package pipe

import (
	"testing"
)

func expect(t *testing.T, result, expected interface{}) {
	if result != expected {
		t.Fatal("Expected:", expected, "\nGot:", result)
	}
}
