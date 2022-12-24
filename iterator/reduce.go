package iterator

// Reduce reduces a iterator down to a single value using a reduction function
func Reduce[S, D any](s Iterator[S], initializer D, f func(D, S) D) D {
	r := initializer
	for s.HasNext() {
		r = f(r, s.Next())
	}
	return r
}
