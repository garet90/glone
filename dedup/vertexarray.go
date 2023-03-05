package dedup

import "github.com/garet90/glone"

type VertexArray struct {
	glone.VertexArray

	arrayBufferBinding        lazyUpdater[glone.Buffer]
	elementArrayBufferBinding lazyUpdater[glone.Buffer]
}

func newVertexArray(v glone.VertexArray) *VertexArray {
	return &VertexArray{
		VertexArray: v,
	}
}

func (V *VertexArray) child() glone.VertexArray {
	if V == nil {
		return nil
	}
	return V.VertexArray
}

func vertexArrayOrNil(v glone.VertexArray) *VertexArray {
	vv, _ := v.(*VertexArray)
	return vv
}
