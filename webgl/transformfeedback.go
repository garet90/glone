package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type TransformFeedback struct {
	value js.Value
}

func transformFeedbackOrNil(v glone.TransformFeedback) any {
	vv, ok := v.(TransformFeedback)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.TransformFeedback = TransformFeedback{}
