package dedup

import "gfx.cafe/ghalliday1/glone"

type Shader struct {
	glone.Shader

	shaderType    glone.Enum
	deleteStatus  bool
	compileStatus bool
}
