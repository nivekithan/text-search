package common

type ReverseIndex[T comparable] struct {
	index map[T]Set[int]
}

func NewReverseIndex[T comparable]() *ReverseIndex[T] {
	return &ReverseIndex[T]{
		index: map[T]Set[int]{},
	}
}

func (r *ReverseIndex[T]) AddEntry(entry T, id int) {
	existingSet, ok := r.index[entry]

	if !ok {
		newSet := NewSet[int]()

		newSet.Add(id)
		r.index[entry] = *newSet
		return
	}

	existingSet.Add(id)
}

func (r *ReverseIndex[T]) GetEntry(entry T) []int {
	existingsIds, ok := r.index[entry]

	if !ok {
		return []int{}
	}

	return existingsIds.Values()
}
