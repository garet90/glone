package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type ShaderPrecisionFormat struct {
	value js.Value
}

func shaderPrecisionFormatOrNil(v glone.ShaderPrecisionFormat) any {
	vv, ok := v.(ShaderPrecisionFormat)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.ShaderPrecisionFormat = ShaderPrecisionFormat{}
