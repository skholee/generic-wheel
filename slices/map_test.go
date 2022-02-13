package slices

import (
	"reflect"
	"testing"
)

func Test_map_should_contains_elems_apply_a_function_to_each_element(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		f      func(string) int
		expect []int
	}{
		{
			name:   "empty",
			input:  []string{},
			f:      func(string) int { return 0 },
			expect: []int{},
		}, {
			name:   "transform to length",
			input:  []string{"1", "12", "123"},
			f:      func(s string) int { return len(s) },
			expect: []int{1, 2, 3},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := Map(tt.input, tt.f)

			if !reflect.DeepEqual(r, tt.expect) {
				t.Errorf("%v should equal %v", r, tt.expect)
			}
		})
	}
}
