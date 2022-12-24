package iterator

func Range(begin, end, step int) Iterator[int] {
	cur := begin
	hasNext := func() bool {
		return (cur >= begin && cur < end) || (cur > end && cur <= begin)
	}
	next := func() int {
		if hasNext() {
			c := cur
			cur = cur + step
			return c
		}
		panic("iterator without next element")
	}
	return NewIterator(hasNext, next)
}
