package sets

func Difference[T comparable](l, r Set[T]) Set[T] {
	re := l.Clone()
	re.Remove(r.ToSlice()...)
	return re
}
