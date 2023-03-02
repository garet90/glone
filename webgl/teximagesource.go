package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type TexImageSource struct {
	value js.Value
}

func (*TexImageSource) GLOneTexImageSource() {}

func MakeTexImageSource(value js.Value) TexImageSource {
	return TexImageSource{
		value,
	}
}

func NewTexImageSource(value js.Value) *TexImageSource {
	v := MakeTexImageSource(value)
	return &v
}

func texImageSourceOrNil(v glone.TexImageSource) any {
	vv, ok := v.(*TexImageSource)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.TexImageSource = (*TexImageSource)(nil)
