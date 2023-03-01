package dedup

import "gfx.cafe/ghalliday1/glone"

type Buffer struct {
	glone.Buffer

	bufferSize  int
	bufferUsage glone.Enum
}
