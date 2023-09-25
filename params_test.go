package moonshine

import "testing"

func TestParams(t *testing.T) {
	params := Params{
		Param{
			Key:   "key",
			Value: "value",
		},
	}
	if params.ByName("key") != "value" {
		t.Error("params.ByName() is not working")
	}
	if _, ok := params.Get("key"); !ok {
		t.Error("params.Get() is not working")
	}
}
