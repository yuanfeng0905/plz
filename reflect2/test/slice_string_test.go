package test

import (
	"testing"
	"github.com/v2pro/plz/reflect2"
)

func Test_slice_string(t *testing.T) {
	t.Run("Set", testOp(func(api reflect2.API) interface{} {
		obj := []string{"hello", "world"}
		valType := api.TypeOf(obj).(reflect2.SliceType)
		valType.Set(&obj, 0, "hi")
		valType.Set(&obj, 1, "there")
		return obj
	}))
}