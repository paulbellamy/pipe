package pipe

import (
	"fmt"
	"testing"
)

func TestMapCatSlice(t *testing.T) {
	printer := func(item int) []string {
		return []string{fmt.Sprintf("%d", item), fmt.Sprintf("%.3d", item)}
	}
	in := []int{7, 4, 5}
	result := MapCatSlice(printer, in).([]string)

	expected := []string{"7", "007", "4", "004", "5", "005"}
	expect(t, len(result), len(expected))
	for i := 0; i < len(result); i++ {
		expect(t, result[i], expected[i])
	}
}
