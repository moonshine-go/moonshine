package moonshine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	params := Params{
		Param{
			Key:   "key",
			Value: "value",
		},
	}
	assert.Equal(t, "value", params.ByName("key"))
	if _, ok := params.Get("key"); !ok {
		t.Error("params.Get() is not working")
	}
}
