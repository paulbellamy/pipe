package pipe

import (
	"testing"
)

func TestFromChan(t *testing.T) {
	source := make(chan int)
	s := New(source)

	go func() {
		source <- 1
		source <- 2
		source <- 3
		close(source)
	}()

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
