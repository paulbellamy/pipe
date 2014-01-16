package pipe

import (
	"testing"
)

func TestFromSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	s := New(slice)
	assertEqual(t, false, s.Empty())

	assertEqual(t, 1, s.First().(int))

	// Check it is immutable
	assertEqual(t, 1, s.First().(int))

	r := s.Rest()
	assertEqual(t, 2, r.First().(int))

	r = r.Rest()
	assertEqual(t, 3, r.First().(int))

	// Check when it's empty
	r = r.Rest()
	assert(t, r.Empty())
}
