package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type Framebuffer struct {
	value js.Value
}

func framebufferOrNil(v glone.Framebuffer) any {
	vv, ok := v.(Framebuffer)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Framebuffer = Framebuffer{}
