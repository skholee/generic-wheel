package slices

// Reduce reduces a slice down to a single value using a reduction function
func Reduce[S, D any](s []S, initializer D, f func(D, S) D) D {
	r := initializer
	for _, e := range s {
		r = f(r, e)
	}
	return r
}
