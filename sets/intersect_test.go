package sets

import (
	"reflect"
	"testing"
)

func Test_set_should_conains_meeting_elems_belong_s1_s2_and_s3_after_intersect(t *testing.T) {
	tests := []struct {
		name string
		sets []Set[string]
		expect Set[string]
	} {
		{
			name: "normal",
			sets: []Set[string]{
				NewSet("elem1", "elem2", "elem3", "elem5", "elem7", "elem9"),
				NewSet("elem1", "elem5", "elem9", "elem11"),
				NewSet("elem2", "elem7", "elem9", "elem18")},
			expect: NewSet("elem9"),
		},
		{
			name: "empty set",
			sets: []Set[string] {
				NewSet("elem1", "elem2"),
				NewSet("elem2", "elem5"),
				NewSet("elem5")},
			expect: NewSet[string](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			
			r := Intersect(tt.sets...)
			if !reflect.DeepEqual(r, tt.expect) {
				t.Errorf("%v should equal %v", r, tt.expect)
			}
		})
	}
}
