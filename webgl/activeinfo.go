package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type ActiveInfo struct {
	value js.Value
}

func (ActiveInfo) GLOneActiveInfo() {}

func activeInfoOrNil(v glone.ActiveInfo) any {
	vv, ok := v.(ActiveInfo)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.ActiveInfo = ActiveInfo{}
