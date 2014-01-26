package pipe

import (
	"testing"
)

func TestZipSlice(t *testing.T) {
	in := []int{5, 10, 20}
	other := []int{6, 11}
	result := ZipSlice(in, other).([][]int)

	expected := [][]int{{5, 6}, {10, 11}}
	if len(result) != len(expected) {
		t.Fatal("expected output to match", expected, "but got", result)
	}

	for i := 0; i < len(expected); i++ {
		if len(result[i]) != len(expected[i]) {
			t.Fatal("expected output to match", expected, "but got", result)
		}

		for j := 0; j < len(result[i]); j++ {
			if result[i][j] != expected[i][j] {
				t.Fatal("expected channel output to match", expected, "but got", result)
			}
		}
	}
}
