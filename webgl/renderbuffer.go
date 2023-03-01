package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type Renderbuffer struct {
	value js.Value
}

func renderbufferOrNil(v glone.Renderbuffer) any {
	vv, ok := v.(Renderbuffer)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Renderbuffer = Renderbuffer{}
