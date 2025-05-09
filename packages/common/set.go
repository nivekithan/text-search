package common

type Set[T comparable] struct {
	value map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{value: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.value[value] = struct{}{}
}

func (s *Set[T]) Values() []T {
	keys := []T{}

	for k := range s.value {
		keys = append(keys, k)
	}

	return keys
}
