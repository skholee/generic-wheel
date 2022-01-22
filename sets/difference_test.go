package sets

import (
	"reflect"
	"testing"
)

func Test_set_should_conains_elems_belong_A_but_not_to_B_after_diff(t *testing.T) {
	tests := []struct {
		name string
		A Set[string]
		B Set[string]
		expect Set[string]
	} {
		{
			name: "normal",
			A: NewSet("elem1", "elem2", "elem3"),
			B: NewSet("elem1", "elem9"),
			expect: NewSet("elem2", "elem3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := Difference(tt.A, tt.B)

			if !reflect.DeepEqual(r, tt.expect) {
				t.Errorf("%v should equal %v", r, tt.expect)
			}
		})
	}
}
