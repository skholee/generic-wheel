package iterator

import (
	"reflect"
	"testing"
)

func Test_filter(t *testing.T) {
	i := Range(1, 9, 1)

	i1 := Filter(i, func(i int) bool { return i%2 == 0 })
	s := i1.ToSlice()
	expect := []int{2, 4, 6, 8}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf("%v should equal %v", expect, s)
	}
}
