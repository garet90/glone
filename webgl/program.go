package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type Program struct {
	value js.Value
}

func (*Program) GLOneProgram() {}

func programOrNil(v glone.Program) any {
	vv, ok := v.(*Program)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Program = (*Program)(nil)
