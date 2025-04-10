package main

type Eventer[T any] struct {
	events map[int]func(T)
}

func (e *Eventer[T]) Add(name int, callback func(T)) {
	if e.events == nil {
		e.events = map[int]func(T){}
	}

	e.events[name] = callback
}

func (e *Eventer[T]) Invoke(name int, data T) {
	for n, f := range e.events {
		if n == name {
			f(data)
		}
	}
}
