package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type ShaderPrecisionFormat struct {
	value js.Value
}

func (*ShaderPrecisionFormat) GLOneShaderPrecisionFormat() {}

func shaderPrecisionFormatOrNil(v glone.ShaderPrecisionFormat) any {
	vv, ok := v.(*ShaderPrecisionFormat)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.ShaderPrecisionFormat = (*ShaderPrecisionFormat)(nil)
