package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type Texture struct {
	value js.Value
}

func (*Texture) GLOneTexture() {}

func textureOrNil(v glone.Texture) any {
	vv, ok := v.(*Texture)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Texture = (*Texture)(nil)
