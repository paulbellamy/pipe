package pipe

import (
	"testing"
)

func TestOntoChan(t *testing.T) {
	in := []int{0, 1, 2, 3, 4}
	out := OntoChan(in).(chan int)

	result := []int{}
	for x := range out {
		result = append(result, x)
	}

	expect(t, len(result), len(in))
	for i := 0; i < len(result); i++ {
		expect(t, result[i], in[i])
	}
}
