package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type BufferSource struct {
	value js.Value
}

func (*BufferSource) GLOneBufferSource() {}

func bufferSourceOrNil(v glone.BufferSource) any {
	vv, ok := v.(*BufferSource)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.BufferSource = (*BufferSource)(nil)
