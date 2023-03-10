package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type Query struct {
	value js.Value
}

func (*Query) GLOneQuery() {}

func queryOrNil(v glone.Query) any {
	vv, ok := v.(*Query)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Query = (*Query)(nil)
