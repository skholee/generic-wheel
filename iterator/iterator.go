package iterator

// Iterator allow to iterate over a sequence of elements
type Iterator[T any] interface {
	// Check if a next element availiable
	HasNext() bool
	// Return the next element and advances the iteraotr
	Next() T
	// Transform to a slice
	ToSlice() []T
}

// is impliment of Iterator
type iterator[T any] struct {
	hasNext func() bool
	next    func() T
}

// The factory of build Iteraotr
func NewIterator[T any](hasNext func() bool, next func() T) Iterator[T] {
	return &iterator[T]{
		hasNext: hasNext,
		next:    next,
	}
}

func FromSlice[T any](s []T) Iterator[T] {
	index := 0
	length := len(s)
	hasNext := func() bool {
		return index < length
	}
	next := func() T {
		cur := s[index]
		index = index + 1
		return cur
	}
	return NewIterator(hasNext, next)
}

func (i *iterator[T]) HasNext() bool {
	return i.hasNext()
}

func (i *iterator[T]) Next() T {
	return i.next()
}

func (i *iterator[T]) ToSlice() []T {
	res := []T{}
	for i.HasNext() {
		res = append(res, i.Next())
	}
	return res
}
