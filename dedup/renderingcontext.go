package dedup

import "gfx.cafe/ghalliday1/glone"

type RenderingContext struct {
	glone.RenderingContext
}

func MakeRenderingContext(context glone.RenderingContext) RenderingContext {
	return RenderingContext{
		context,
	}
}

func NewRenderingContext(context glone.RenderingContext) *RenderingContext {
	rc := MakeRenderingContext(context)
	return &rc
}
