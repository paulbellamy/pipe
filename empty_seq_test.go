package pipe

import (
	"testing"
)

func TestEmptySeq(t *testing.T) {
	assertEqual(t, nil, Empty.First())
	assertEqual(t, Empty, Empty.Rest())
	assert(t, Empty.Empty())
}
