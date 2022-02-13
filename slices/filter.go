package slices

// Filter filters values from a slice using a predicate
// return a new slice contains elements of this slice which satisfy a predicate
func Filter[T any](s []T, predicate func(T) bool) []T {
	r := []T{}
	for _, e := range s {
		if predicate(e) {
			r = append(r, e)
		}
	}
	return r
}
