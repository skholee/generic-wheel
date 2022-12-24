package iterator

import "github.com/skholee/generic-wheel/option"

// Filter filters values from a slice using a predicate
// return a new slice contains elements of this slice which satisfy a predicate
func Filter[T any](i Iterator[T], predicate func(T) bool) Iterator[T] {
	nextOption := func() option.Option[T] {
		for i.HasNext() {
			cur := i.Next()
			if predicate(cur) {
				return option.Some(cur)
			}
		}
		return option.None[T]()
	}
	var peek option.Option[T]
	hasNext := func() bool {
		if peek == nil {
			peek = nextOption()
		}
		return !peek.IsEmpty()
	}
	next := func() T {
		if peek == nil {
			peek = nextOption()
		}
		t := peek.Get()
		peek = nil
		return t
	}
	return NewIterator(hasNext, next)
}
