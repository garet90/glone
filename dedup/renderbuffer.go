package dedup

import "gfx.cafe/ghalliday1/glone"

type Renderbuffer struct {
	glone.Renderbuffer

	width          int
	height         int
	internalFormat glone.Enum
	redSize        int
	greenSize      int
	blueSize       int
	alphaSize      int
	depthSize      int
	samples        int
	stencilSize    int
}
