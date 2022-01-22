package sets

import (
	"reflect"
	"testing"
)

func Test_set_should_conains_either_s1_s2_or_s3_set_elems_after_union(t *testing.T) {
	s1 := NewSet("elem1", "elem2")
	s2 := NewSet("elem2", "elem6")
	s3 := NewSet("elem9")

	r := Union(s1, s2, s3)

	expect := NewSet("elem1", "elem2", "elem6", "elem9")

	if !reflect.DeepEqual(r, expect) {
		t.Errorf("%v should equal %v", r, expect)
	}
}
