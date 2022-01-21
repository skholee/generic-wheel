package dictionary

import "testing"

func Test_dict_should_empty_after_init(t *testing.T) {
	dict := Dictionary[string, int]{}
	if len(dict) != 0 {
		t.Errorf("%d should equal 0", len(dict))
	}
}

func Test_dict_should_able_to_index_key_after_append(t *testing.T) {
	dict := Dictionary[string, int]{}
	dict["key"] = 1

	if dict["key"] != 1 {
		t.Errorf("%d should equal 1", dict["key"])
	}
}

func Test_dict_should_index_last_value_after_multiple_updates(t *testing.T) {
	dict := Dictionary[string, int]{}
	dict["key"] = 1
	dict["key"] = 2
	dict["key"] = 3

	if dict["key"] != 3 {
		t.Errorf("%d should equal 3", dict["key"])
	}
}

func Test_dict_should_cannot_able_to_index_key_after_delete(t *testing.T) {
	dict := Dictionary[string, int]{"key": 1}
	
	if len(dict) != 1 {
		t.Errorf("%d should equal 1", len(dict))
	}

	delete(dict, "key")

	if _, ok := dict["key"]; ok != false {
		t.Errorf("%t should equal false", ok)
	}
}
