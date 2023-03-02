package webgl

import (
	"github.com/garet90/glone"
	"syscall/js"
)

type Sampler struct {
	value js.Value
}

func (*Sampler) GLOneSampler() {}

func samplerOrNil(v glone.Sampler) any {
	vv, ok := v.(*Sampler)
	if !ok {
		return nil
	}
	return vv.value
}

var _ glone.Sampler = (*Sampler)(nil)
