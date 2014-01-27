package pipe

import (
	"fmt"
)

func ExampleChain() {
	type stooge struct {
		name string
		age  int
	}

	stooges := []stooge{{"curly", 25}, {"moe", 21}, {"larry", 23}}

	youngest := Chain(stooges).
		SortBy(FieldByName("age")).
		Map(func(s stooge) string {
		return fmt.Sprintf("%s is %d", s.name, s.age)
	}).
		First().(string)

	fmt.Println(youngest)
}
