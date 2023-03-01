package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type Buffer struct {
	value js.Value
}

func bufferOrNil(v glone.Buffer) any {
	vv, ok := v.(Buffer)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Buffer = Buffer{}
