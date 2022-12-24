package iterator

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_map(t *testing.T) {
	i := FromSlice([]int{1, 2})

	i1 := Map(i, func(i int) string { return strconv.Itoa(i) })
	s := i1.ToSlice()
	expect := []string{"1", "2"}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf("%v should equal %v", expect, s)
	}
}
