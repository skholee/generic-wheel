package sets

import "testing"

func Test_set_should_empty_after_init(t *testing.T) {
	s := NewSet[string]()

	if s.Len() != 0 {
		t.Errorf("%d should equal 0", s.Len())
	}

	if !s.IsEmpty() {
		t.Errorf("%t should equal true", s.IsEmpty())
	}
}

func Test_set_should_contains_elem_after_added(t *testing.T) {
	s := NewSet[string]()

	s.Add("elem1", "elem2")

	if !s.Contains("elem1") {
		t.Errorf("%t should equal true", s.Contains("elem1"))
	}

	if s.Contains("elem3") {
		t.Errorf("%t should equal false", s.Contains("elem3"))
	}
}

func Test_set_should_not_contains_elem_after_removed(t *testing.T) {
	s := NewSet("elem1", "elem2")

	s.Remove("elem2", "elem3")

	if s.Contains("elem2") {
		t.Errorf("%t should equal false", s.Contains("elem2"))
	}

	if !s.Contains("elem1") {
		t.Errorf("%t should equal true", s.Contains("elem1"))
	}
}
