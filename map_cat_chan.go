package pipe

// MapCatChan is of type: func(fn func(T) []U, input chan T) chan U.
// It returns a chan which receives fn(item) for each item in input.
func MapCatChan(fn, input interface{}) interface{} {
	return FlattenChan(MapChan(fn, input))
}
