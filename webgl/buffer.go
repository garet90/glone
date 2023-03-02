package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type Buffer struct {
	value js.Value
}

func (*Buffer) GLOneBuffer() {}

func bufferOrNil(v glone.Buffer) any {
	vv, ok := v.(*Buffer)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Buffer = (*Buffer)(nil)
