package slices

import (
	"reflect"
	"testing"
)

func Test_reduce_should_should_reduce_a_slice_down_to_a_single_value(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		initializer int
		f           func(int, string) int
		expect      int
	}{
		{
			name:        "empty",
			input:       []string{},
			initializer: 0,
			f:           func(int, string) int { return 1 },
			expect:      0,
		}, {
			name:        "totoal char number",
			input:       []string{"1", "12", "123"},
			initializer: 0,
			f:           func(l int, s string) int { return l + len(s) },
			expect:      6,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := Reduce(tt.input, tt.initializer, tt.f)

			if !reflect.DeepEqual(r, tt.expect) {
				t.Errorf("%v should equal %v", r, tt.expect)
			}
		})
	}
}
