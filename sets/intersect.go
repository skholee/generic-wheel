package sets

// Intersect the intersection of some sets(eg. A, B)
// is the set containing all elements of A that also belong to B
func Intersect[T comparable](sets ...Set[T]) Set[T] {
	records := map[T]int{}

	for _, s := range sets {
		for _, e := range s.ToSlice() {
			records[e] = records[e] + 1
		}
	}

	r := NewSet[T]()

	for e, num := range records {
		if num == len(sets) {
			r.Add(e)
		}
	}

	return r
}
