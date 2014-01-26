package pipe

import (
	"testing"
)

func TestReverseSlice(t *testing.T) {
	in := []int{7, 6, 5}
	result := ReverseSlice(in).([]int)

	expected := []int{5, 6, 7}

	expect(t, len(result), len(expected))
	for i := 0; i < 3; i++ {
		expect(t, result[i], expected[i])
	}

	/*
		expect(t, len(in), len(expected))
		for i := 0; i < 3; i++ {
			expect(t, in[i], expected[i])
		}
	*/
}
