package pipe

import (
	"testing"
)

type methodByNameCustomer struct {
	Name string
}

func (c methodByNameCustomer) String() string {
	return c.Name
}

func TestMethodByName(t *testing.T) {

	alice := methodByNameCustomer{"Alice"}

	getter := MethodByName("String").(func(interface{}) interface{})
	result := getter(alice).(func() string)()
	if result != "Alice" {
		t.Fatal("Expected getter to return 'Alice', but returned,", result)
	}
}
