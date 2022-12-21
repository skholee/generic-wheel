package option

import "testing"

func Test_None(t *testing.T) {
	o := None[string]()

	if !o.IsEmpty() {
		t.Errorf("%t should equal true", o.IsEmpty())
	}

	if o.GetOrElse("default") != "default" {
		t.Errorf("%s should equal default", o.GetOrElse("default"))
	}

	if o.OrElse(Some("some")).Get() != "some" {
		t.Errorf("%s should equal some", o.OrElse(Some("some")).Get())
	}
}

func Test_Some(t *testing.T) {
	o := Some("some")

	if o.IsEmpty() {
		t.Errorf("%t should equal false", o.IsEmpty())
	}

	if o.GetOrElse("default") != "some" {
		t.Errorf("%s should equal some", o.GetOrElse("default"))
	}

	if o.Get() != "some" {
		t.Errorf("%s should equal some", o.Get())
	}

	if o.OrElse(Some("x")).Get() != "some" {
		t.Errorf("%s should equal some", o.OrElse(Some("x")).Get())
	}
}
