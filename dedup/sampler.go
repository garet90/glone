package dedup

import "gfx.cafe/ghalliday1/glone"

type Sampler struct {
	glone.Sampler

	compareFunc glone.Enum
	compareMode glone.Enum
	magFilter   glone.Enum
	maxLod      float32
	minFilter   glone.Enum
	minLod      float32
	wrapR       glone.Enum
	wrapS       glone.Enum
	wrapT       glone.Enum
}
