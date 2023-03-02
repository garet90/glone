package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type TransformFeedback struct {
	value js.Value
}

func (*TransformFeedback) GLOneTransformFeedback() {}

func transformFeedbackOrNil(v glone.TransformFeedback) any {
	vv, ok := v.(*TransformFeedback)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.TransformFeedback = (*TransformFeedback)(nil)
