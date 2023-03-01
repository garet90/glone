package dedup

import "gfx.cafe/ghalliday1/glone"

type Program struct {
	glone.Program

	deleteStatus                bool
	linkStatus                  bool
	validateStatus              bool
	attachedShaders             int
	activeAttributes            int
	activeUniforms              int
	transformFeedbackBufferMode glone.Buffer
	transformFeedbackVaryings   int
	activeUniformBlocks         int
}
