package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type UniformLocation struct {
	value js.Value
}

func (*UniformLocation) GLOneUniformLocation() {}

func uniformLocationOrNil(v glone.UniformLocation) any {
	vv, ok := v.(*UniformLocation)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.UniformLocation = (*UniformLocation)(nil)
