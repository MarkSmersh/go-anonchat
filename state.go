package main

type State[K string | int, V any] struct {
	states map[K]V
}

func (s *State[K, V]) Set(k K, v V) {
	s.states[k] = v
}

func (s *State[K, V]) Get(k K) V {
	return s.states[k]
}
