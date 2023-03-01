package dedup

import "gfx.cafe/ghalliday1/glone"

type Query struct {
	glone.Query

	queryResult          uint
	queryResultAvailable bool
}
