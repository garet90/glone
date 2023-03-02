package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type Framebuffer struct {
	value js.Value
}

func (*Framebuffer) GLOneFramebuffer() {}

func framebufferOrNil(v glone.Framebuffer) any {
	vv, ok := v.(*Framebuffer)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Framebuffer = (*Framebuffer)(nil)
