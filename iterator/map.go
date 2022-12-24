package iterator

// Map used to apply a function to each element in iterator
// return a new iterator contains the transformed elements
func Map[S, D any](s Iterator[S], f func(S) D) Iterator[D] {
	next := func() D {
		return f(s.Next())
	}
	return NewIterator(s.HasNext, next)
}
