package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type UniformLocation struct {
	value js.Value
}

func uniformLocationOrNil(v glone.UniformLocation) any {
	vv, ok := v.(UniformLocation)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.UniformLocation = UniformLocation{}
