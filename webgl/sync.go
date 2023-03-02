package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type Sync struct {
	value js.Value
}

func (*Sync) GLOneSync() {}

func syncOrNil(v glone.Sync) any {
	vv, ok := v.(*Sync)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Sync = (*Sync)(nil)
