package dedup

import "gfx.cafe/ghalliday1/glone"

type Sync struct {
	glone.Sync

	objectType    glone.Enum
	syncStatus    glone.Enum
	syncCondition glone.Enum
	syncFlags     glone.Enum
}
