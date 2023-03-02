package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type VertexArray struct {
	value js.Value
}

func (*VertexArray) GLOneVertexArray() {}

func vertexArrayOrNil(v glone.VertexArray) any {
	vv, ok := v.(*VertexArray)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.VertexArray = (*VertexArray)(nil)
