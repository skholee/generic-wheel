package slices

import (
	"reflect"
	"testing"
)

func Test_filter_should_contains_elems_satisfying_predicate(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		predicate func(string) bool
		expect    []string
	}{
		{
			name:      "empty",
			input:     []string{},
			predicate: func(string) bool { return true },
			expect:    []string{},
		}, {
			name:      "all",
			input:     []string{"1", "11", "111", "1111"},
			predicate: func(string) bool { return true },
			expect:    []string{"1", "11", "111", "1111"},
		}, {
			name:      "filter len less or equal to 2",
			input:     []string{"1", "11", "111", "1111"},
			predicate: func(s string) bool { return len(s) <= 2 },
			expect:    []string{"1", "11"},
		}, {
			name:      "filter nothing",
			input:     []string{"1", "11", "111", "1111"},
			predicate: func(string) bool { return false },
			expect:    []string{},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := Filter(tt.input, tt.predicate)

			if !reflect.DeepEqual(r, tt.expect) {
				t.Errorf("%v should equal %v", r, tt.expect)
			}
		})
	}
}
