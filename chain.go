package pipe

import ()

type chain struct {
	source Seq
}

func Chain(source interface{}) *chain {
	return &chain{New(source)}
}

func (c *chain) SortBy(accessor interface{}) *chain {
	// TODO: Implement this?
	return c
}

func (c *chain) Map(fn interface{}) *chain {
	// TODO: Implement this?
	return c
}

func (c *chain) MapCat(fn interface{}) *chain {
	// TODO: Implement this?
	return c
}

func (c *chain) First() interface{} {
	return c.source.First()
}

func (c *chain) Chan() interface{} {
	result := make(chan string)
	go func() {
		close(result)
	}()
	return result
}
