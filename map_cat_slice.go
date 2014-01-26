package pipe

// MapCatSlice is of type: func(fn func(T) []U, input []T) []U.
// It returns a slice with fn(item) for each item in input.
func MapCatSlice(fn, input interface{}) interface{} {
	return FlattenSlice(MapSlice(fn, input))
}
