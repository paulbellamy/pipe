package pipe

import (
	"math/rand"
	"testing"
)

func TestCallMethodByName(t *testing.T) {

	generator := rand.New(rand.NewSource(1))

	caller := CallMethodByName("Intn", 100).(func(interface{}) []interface{})
	result := caller(generator)[0].(int)
	if result != 0 {
		t.Fatal("Expected getter to return 'Alice', but returned,", result)
	}
}
