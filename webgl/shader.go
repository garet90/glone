package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type Shader struct {
	value js.Value
}

func shaderOrNil(v glone.Shader) any {
	vv, ok := v.(Shader)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Shader = Shader{}
