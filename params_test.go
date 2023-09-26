package moonshine

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestParams(t *testing.T) {
	params := Params{
		Param{
			Key:   "key",
			Value: "value",
		},
	}
	assert.Equal(t, params.ByName("key"), "value")
	_, ok := params.Get("key")
	assert.Equal(t, ok, true)
}
