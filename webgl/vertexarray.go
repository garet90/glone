package webgl

import (
	"gfx.cafe/ghalliday1/glone"
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
