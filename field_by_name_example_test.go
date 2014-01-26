package pipe

import (
	"fmt"
)

func ExampleFieldByName() {
	type customer struct {
		Name string
		Age  int
	}

	customers := []customer{{"Alice", 26}, {"Bob", 46}, {"Yousef", 37}}

	names := MapSlice(FieldByName("Name"), customers)
	ages := MapSlice(FieldByName("Age"), customers)

	fmt.Println(names)
	fmt.Println(ages)

	// Output:
	// [Alice Bob Yousef]
	// [26 46 37]
}
