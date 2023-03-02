package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type Shader struct {
	value js.Value
}

func (*Shader) GLOneShader() {}

func shaderOrNil(v glone.Shader) any {
	vv, ok := v.(*Shader)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Shader = (*Shader)(nil)
