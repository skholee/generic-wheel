package sets

import (
	"reflect"
	"testing"
)

func Test_set_should_conains_either_s1_s2_or_s3_set_elems_after_union(t *testing.T) {
	tests := []struct{
		name string
		sets []Set[string]
		expect Set[string]
	} {
		{
			name: "normal",
			sets: []Set[string]{
				NewSet("elem1", "elem2"),
				NewSet("elem2", "elem6"),
				NewSet("elem9")},
			expect: NewSet("elem1", "elem2", "elem6", "elem9"),
		},
		{
			name: "empty set",
			sets: []Set[string]{
				NewSet[string](),
				NewSet("elem2", "elem6"),
				NewSet("elem9")},
			expect: NewSet("elem2", "elem6", "elem9"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			t.Parallel()

			r := Union(tt.sets...)

			if !reflect.DeepEqual(r, tt.expect) {
				t.Errorf("%v should equal %v", r, tt.expect)
			}
		})
	}
}
