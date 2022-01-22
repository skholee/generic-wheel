package sets

func Union[T comparable](sets ...Set[T]) Set[T] {
	r := NewSet[T]()
	for _, s := range sets {
		r.Add(s.ToSlice()...)
	}
	return r
}