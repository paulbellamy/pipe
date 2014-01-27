package pipe

import (
	"fmt"
)

type friend struct {
	Name string
	Age  int
}

func (c friend) String() string {
	return fmt.Sprintf("%s is %d", c.Name, c.Age)
}

func ExampleMethodByName() {

	friends := []friend{{"Alice", 26}, {"Bob", 46}, {"Yousef", 37}}

	formatters := MapSlice(MethodByName("String"), friends).([]interface{})

	for _, formatter := range formatters {
		fmt.Println(formatter.(func() string)())
	}

	// Output:
	// Alice is 26
	// Bob is 46
	// Yousef is 37
}
