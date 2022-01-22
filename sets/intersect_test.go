package sets

import (
	"reflect"
	"testing"
)

func Test_set_should_conains_meeting_elems_belong_s1_s2_and_s3_after_intersect(t *testing.T) {
	s1 := NewSet("elem1", "elem2", "elem3", "elem5", "elem7", "elem9")
	s2 := NewSet("elem1", "elem5", "elem9", "elem11")
	s3 := NewSet("elem2", "elem7", "elem9", "elem18")

	r := Intersect(s1, s2, s3)

	expect := NewSet("elem9")

	if !reflect.DeepEqual(r, expect) {
		t.Errorf("%v should equal %v", r, expect)
	}
}
