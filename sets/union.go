package sets

// Union the union of some sets(eg. A, B)
// is the set of elements that are in either A or B
func Union[T comparable](sets ...Set[T]) Set[T] {
	r := NewSet[T]()
	for _, s := range sets {
		r.Add(s.ToSlice()...)
	}
	return r
}
