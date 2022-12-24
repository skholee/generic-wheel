package iterator

import (
	"strconv"
	"testing"
)

func Test_reduce(t *testing.T) {
	i := Range(1, 3, 1)

	r := Reduce(i, "", func(s string, i int) string { return s + strconv.Itoa(i) })
	if r != "12" {
		t.Errorf("%v should equal %v", r, "12")
	}
}
