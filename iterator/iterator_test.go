package iterator

import (
	"reflect"
	"testing"
)

func Test_iterator(t *testing.T) {
	i := Range(1, 3, 1)

	s := i.ToSlice()

	expect := []int{1, 2}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf("%v should equal %v", expect, s)
	}
}
