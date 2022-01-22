package sets

// Difference the difference of two sets(eg. A, B)
// is the set of elements which are in A but not in B
func Difference[T comparable](l, r Set[T]) Set[T] {
	re := l.Clone()
	re.Remove(r.ToSlice()...)
	return re
}
