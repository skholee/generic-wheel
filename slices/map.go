package slices

// Map used to apply a function to each element in slice
// return a new slice contains the transformed elements
func Map[S, D any](s []S, f func(S) D) []D {
	r := make([]D, 0, len(s))
	for _, e := range s {
		r = append(r, f(e))
	}
	return r
}
