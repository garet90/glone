package dedup

import "gfx.cafe/ghalliday1/glone"

type Texture struct {
	glone.Texture

	baseLevel       int
	compareFunc     glone.Enum
	compareMode     glone.Enum
	immutableFormat bool
	immutableLevels uint
	magFilter       glone.Enum
	maxLevel        int
	maxLod          float32
	minFilter       glone.Enum
	minLod          float32
	wrapR           glone.Enum
	wrapS           glone.Enum
	wrapT           glone.Enum
}
