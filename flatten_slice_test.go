package pipe

import (
	"testing"
)

func TestFlattenSlice(t *testing.T) {
	in := [][]int{{1, 2}, {3, 4}, {5, 6}}
	result := FlattenSlice(in).([]int)

	expected := []int{1, 2, 3, 4, 5, 6}

	if len(expected) != len(result) {
		t.Fatal("expected output to have 6 elements, but there were", len(result))
	}

	for i := 0; i < len(expected); i++ {
		expect(t, result[i], expected[i])
	}
}

func TestFlattenSliceWhenAlreadyFlat(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	result := FlattenSlice(in).([]int)

	expected := []int{1, 2, 3, 4, 5, 6}

	if len(expected) != len(result) {
		t.Fatal("expected output to have 6 elements, but there were", len(result))
	}

	for i := 0; i < len(expected); i++ {
		expect(t, result[i], expected[i])
	}
}
