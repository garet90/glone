package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type Query struct {
	value js.Value
}

func queryOrNil(v glone.Query) any {
	vv, ok := v.(Query)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Query = Query{}
