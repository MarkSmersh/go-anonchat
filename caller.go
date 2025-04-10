package main

type Caller[T any] struct {
	calls []func(T)
}

func (c *Caller[T]) Add(callback func(T)) {
	c.calls = append(c.calls, callback)
}

func (c *Caller[T]) Invoke(data T) {
	for i := range c.calls {
		c.calls[i](data)
	}
}
