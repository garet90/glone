package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type TexImageSource struct {
	value js.Value
}

func (*TexImageSource) GLOneTexImageSource() {}

func texImageSourceOrNil(v glone.TexImageSource) any {
	vv, ok := v.(*TexImageSource)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.TexImageSource = (*TexImageSource)(nil)
