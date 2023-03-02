package dedup

import (
	"gfx.cafe/ghalliday1/glone"
)

type lazyUpdater[T comparable] struct {
	target  T
	current T
}

func (L *lazyUpdater[T]) isCurrent() bool {
	return L.target == L.current
}

func (L *lazyUpdater[T]) makeCurrent() {
	L.current = L.target
}

type RenderingContext struct {
	glone.RenderingContext

	arrayBufferBinding             lazyUpdater[glone.Buffer]
	elementArrayBufferBinding      lazyUpdater[glone.Buffer]
	copyReadBufferBinding          lazyUpdater[glone.Buffer]
	copyWriteBufferBinding         lazyUpdater[glone.Buffer]
	transformFeedbackBufferBinding lazyUpdater[glone.Buffer]
	uniformBufferBinding           lazyUpdater[glone.Buffer]
	pixelPackBufferBinding         lazyUpdater[glone.Buffer]
	pixelUnpackBufferBinding       lazyUpdater[glone.Buffer]
	drawFramebufferBinding         lazyUpdater[glone.Framebuffer]
	readFramebufferBinding         lazyUpdater[glone.Framebuffer]
	renderbufferBinding            lazyUpdater[glone.Renderbuffer]
	vertexArrayBinding             lazyUpdater[glone.VertexArray]
	programBinding                 lazyUpdater[glone.Program]
}

func MakeRenderingContext(context glone.RenderingContext) RenderingContext {
	return RenderingContext{
		RenderingContext: context,
	}
}

func NewRenderingContext(context glone.RenderingContext) *RenderingContext {
	rc := MakeRenderingContext(context)
	return &rc
}

func (R *RenderingContext) requiresArrayBuffer() {
	if R.arrayBufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindBuffer(glone.ARRAY_BUFFER, R.arrayBufferBinding.target)
	R.arrayBufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresElementArrayBuffer() {
	if R.elementArrayBufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindBuffer(glone.ELEMENT_ARRAY_BUFFER, R.elementArrayBufferBinding.target)
	R.elementArrayBufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresCopyReadBuffer() {
	if R.copyReadBufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindBuffer(glone.COPY_READ_BUFFER, R.copyReadBufferBinding.target)
	R.copyReadBufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresCopyWriteBuffer() {
	if R.copyWriteBufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindBuffer(glone.COPY_WRITE_BUFFER, R.copyWriteBufferBinding.target)
	R.copyWriteBufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresTransformFeedbackBuffer() {
	if R.transformFeedbackBufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindBuffer(glone.TRANSFORM_FEEDBACK_BUFFER, R.transformFeedbackBufferBinding.target)
	R.transformFeedbackBufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresUniformBuffer() {
	if R.uniformBufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindBuffer(glone.UNIFORM_BUFFER, R.uniformBufferBinding.target)
	R.uniformBufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresPixelPackBuffer() {
	if R.pixelPackBufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindBuffer(glone.PIXEL_PACK_BUFFER, R.pixelPackBufferBinding.target)
	R.pixelPackBufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresPixelUnpackBuffer() {
	if R.pixelUnpackBufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindBuffer(glone.PIXEL_UNPACK_BUFFER, R.pixelUnpackBufferBinding.target)
	R.pixelUnpackBufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresBuffer(target glone.Enum) {
	switch target {
	case glone.ARRAY_BUFFER:
		R.requiresArrayBuffer()
	case glone.ELEMENT_ARRAY_BUFFER:
		R.requiresElementArrayBuffer()
	case glone.COPY_READ_BUFFER:
		R.requiresCopyReadBuffer()
	case glone.COPY_WRITE_BUFFER:
		R.requiresCopyWriteBuffer()
	case glone.TRANSFORM_FEEDBACK_BUFFER:
		R.requiresTransformFeedbackBuffer()
	case glone.UNIFORM_BUFFER:
		R.requiresUniformBuffer()
	case glone.PIXEL_PACK_BUFFER:
		R.requiresPixelPackBuffer()
	case glone.PIXEL_UNPACK_BUFFER:
		R.requiresPixelUnpackBuffer()
	}
}

func (R *RenderingContext) BindBuffer(target glone.Enum, buffer glone.Buffer) {
	switch target {
	case glone.ARRAY_BUFFER:
		R.arrayBufferBinding.target = buffer
	case glone.ELEMENT_ARRAY_BUFFER:
		R.elementArrayBufferBinding.target = buffer
	case glone.COPY_READ_BUFFER:
		R.copyReadBufferBinding.target = buffer
	case glone.COPY_WRITE_BUFFER:
		R.copyWriteBufferBinding.target = buffer
	case glone.TRANSFORM_FEEDBACK_BUFFER:
		R.transformFeedbackBufferBinding.target = buffer
	case glone.UNIFORM_BUFFER:
		R.uniformBufferBinding.target = buffer
	case glone.PIXEL_PACK_BUFFER:
		R.pixelPackBufferBinding.target = buffer
	case glone.PIXEL_UNPACK_BUFFER:
		R.pixelUnpackBufferBinding.target = buffer
	}
}

func (R *RenderingContext) requiresReadFramebuffer() {
	if R.readFramebufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindFramebuffer(glone.READ_FRAMEBUFFER, R.readFramebufferBinding.target)
	R.readFramebufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresDrawFramebuffer() {
	if R.drawFramebufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindFramebuffer(glone.DRAW_FRAMEBUFFER, R.drawFramebufferBinding.target)
	R.drawFramebufferBinding.makeCurrent()
}

func (R *RenderingContext) BindFramebuffer(target glone.Enum, framebuffer glone.Framebuffer) {
	switch target {
	case glone.FRAMEBUFFER:
		R.drawFramebufferBinding.target = framebuffer
		R.readFramebufferBinding.target = framebuffer
	case glone.READ_FRAMEBUFFER:
		R.readFramebufferBinding.target = framebuffer
	case glone.DRAW_FRAMEBUFFER:
		R.drawFramebufferBinding.target = framebuffer
	}
}

func (R *RenderingContext) requiresRenderbuffer() {
	if R.renderbufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindRenderbuffer(glone.RENDERBUFFER, R.renderbufferBinding.target)
	R.renderbufferBinding.makeCurrent()
}

func (R *RenderingContext) BindRenderbuffer(target glone.Enum, renderbuffer glone.Renderbuffer) {
	switch target {
	case glone.RENDERBUFFER:
		R.renderbufferBinding.target = renderbuffer
	}
}

func (R *RenderingContext) requiresVertexArray() {
	if R.vertexArrayBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindVertexArray(R.vertexArrayBinding.target)
	R.vertexArrayBinding.makeCurrent()
}

func (R *RenderingContext) BindVertexArray(array glone.VertexArray) {
	R.vertexArrayBinding.target = array
}

func (R *RenderingContext) requiresProgram() {
	if R.programBinding.isCurrent() {
		return
	}
	R.RenderingContext.UseProgram(R.programBinding.target)
	R.programBinding.makeCurrent()
}

func (R *RenderingContext) UseProgram(program glone.Program) {
	R.programBinding.target = program
}

func (R *RenderingContext) Uniform1f(location glone.UniformLocation, x float32) {
	R.requiresProgram()
	R.RenderingContext.Uniform1f(location, x)
}

func (R *RenderingContext) Uniform2f(location glone.UniformLocation, x, y float32) {
	R.requiresProgram()
	R.RenderingContext.Uniform2f(location, x, y)
}

func (R *RenderingContext) Uniform3f(location glone.UniformLocation, x, y, z float32) {
	R.requiresProgram()
	R.RenderingContext.Uniform3f(location, x, y, z)
}

func (R *RenderingContext) Uniform4f(location glone.UniformLocation, x, y, z, w float32) {
	R.requiresProgram()
	R.RenderingContext.Uniform4f(location, x, y, z, w)
}

func (R *RenderingContext) Uniform1i(location glone.UniformLocation, x int32) {
	R.requiresProgram()
	R.RenderingContext.Uniform1i(location, x)
}

func (R *RenderingContext) Uniform2i(location glone.UniformLocation, x, y int32) {
	R.requiresProgram()
	R.RenderingContext.Uniform2i(location, x, y)
}

func (R *RenderingContext) Uniform3i(location glone.UniformLocation, x, y, z int32) {
	R.requiresProgram()
	R.RenderingContext.Uniform3i(location, x, y, z)
}

func (R *RenderingContext) Uniform4i(location glone.UniformLocation, x, y, z, w int32) {
	R.requiresProgram()
	R.RenderingContext.Uniform4i(location, x, y, z, w)
}

func (R *RenderingContext) Uniform1ui(location glone.UniformLocation, v0 uint32) {
	R.requiresProgram()
	R.RenderingContext.Uniform1ui(location, v0)
}

func (R *RenderingContext) Uniform2ui(location glone.UniformLocation, v0, v1 uint32) {
	R.requiresProgram()
	R.RenderingContext.Uniform2ui(location, v0, v1)
}

func (R *RenderingContext) Uniform3ui(location glone.UniformLocation, v0, v1, v2 uint32) {
	R.requiresProgram()
	R.RenderingContext.Uniform3ui(location, v0, v1, v2)
}

func (R *RenderingContext) Uniform4ui(location glone.UniformLocation, v0, v1, v2, v3 uint32) {
	R.requiresProgram()
	R.RenderingContext.Uniform4ui(location, v0, v1, v2, v3)
}

func (R *RenderingContext) Uniform1fv(location glone.UniformLocation, v []float32) {
	R.requiresProgram()
	R.RenderingContext.Uniform1fv(location, v)
}

func (R *RenderingContext) Uniform2fv(location glone.UniformLocation, v []float32) {
	R.requiresProgram()
	R.RenderingContext.Uniform2fv(location, v)
}

func (R *RenderingContext) Uniform3fv(location glone.UniformLocation, v []float32) {
	R.requiresProgram()
	R.RenderingContext.Uniform3fv(location, v)
}

func (R *RenderingContext) Uniform4fv(location glone.UniformLocation, v []float32) {
	R.requiresProgram()
	R.RenderingContext.Uniform4fv(location, v)
}

func (R *RenderingContext) Uniform1iv(location glone.UniformLocation, v []int32) {
	R.requiresProgram()
	R.RenderingContext.Uniform1iv(location, v)
}

func (R *RenderingContext) Uniform2iv(location glone.UniformLocation, v []int32) {
	R.requiresProgram()
	R.RenderingContext.Uniform2iv(location, v)
}

func (R *RenderingContext) Uniform3iv(location glone.UniformLocation, v []int32) {
	R.requiresProgram()
	R.RenderingContext.Uniform3iv(location, v)
}

func (R *RenderingContext) Uniform4iv(location glone.UniformLocation, v []int32) {
	R.requiresProgram()
	R.RenderingContext.Uniform4iv(location, v)
}

func (R *RenderingContext) Uniform1uiv(location glone.UniformLocation, data []uint32) {
	R.requiresProgram()
	R.RenderingContext.Uniform1uiv(location, data)
}

func (R *RenderingContext) Uniform2uiv(location glone.UniformLocation, data []uint32) {
	R.requiresProgram()
	R.RenderingContext.Uniform2uiv(location, data)
}

func (R *RenderingContext) Uniform3uiv(location glone.UniformLocation, data []uint32) {
	R.requiresProgram()
	R.RenderingContext.Uniform3uiv(location, data)
}

func (R *RenderingContext) Uniform4uiv(location glone.UniformLocation, data []uint32) {
	R.requiresProgram()
	R.RenderingContext.Uniform4uiv(location, data)
}

func (R *RenderingContext) UniformMatrix2fv(location glone.UniformLocation, transpose bool, value []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix2fv(location, transpose, value)
}

func (R *RenderingContext) UniformMatrix3fv(location glone.UniformLocation, transpose bool, value []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix3fv(location, transpose, value)
}

func (R *RenderingContext) UniformMatrix4fv(location glone.UniformLocation, transpose bool, value []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix4fv(location, transpose, value)
}

func (R *RenderingContext) UniformMatrix3x2fv(location glone.UniformLocation, transpose bool, data []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix3x2fv(location, transpose, data)
}

func (R *RenderingContext) UniformMatrix4x2fv(location glone.UniformLocation, transpose bool, data []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix4x2fv(location, transpose, data)
}

func (R *RenderingContext) UniformMatrix2x3fv(location glone.UniformLocation, transpose bool, data []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix2x3fv(location, transpose, data)
}

func (R *RenderingContext) UniformMatrix4x3fv(location glone.UniformLocation, transpose bool, data []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix4x3fv(location, transpose, data)
}

func (R *RenderingContext) UniformMatrix2x4fv(location glone.UniformLocation, transpose bool, data []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix2x4fv(location, transpose, data)
}

func (R *RenderingContext) UniformMatrix3x4fv(location glone.UniformLocation, transpose bool, data []float32) {
	R.requiresProgram()
	R.RenderingContext.UniformMatrix3x4fv(location, transpose, data)
}

func (R *RenderingContext) EnableVertexAttribArray(index uint32) {
	R.requiresVertexArray()
	R.RenderingContext.EnableVertexAttribArray(index)
}

func (R *RenderingContext) DisableVertexAttribArray(index uint32) {
	R.requiresVertexArray()
	R.RenderingContext.DisableVertexAttribArray(index)
}

func (R *RenderingContext) VertexAttrib1f(index uint32, x float32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttrib1f(index, x)
}

func (R *RenderingContext) VertexAttrib2f(index uint32, x, y float32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttrib2f(index, x, y)
}

func (R *RenderingContext) VertexAttrib3f(index uint32, x, y, z float32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttrib3f(index, x, y, z)
}

func (R *RenderingContext) VertexAttrib4f(index uint32, x, y, z, w float32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttrib4f(index, x, y, z, w)
}

func (R *RenderingContext) VertexAttrib1fv(index uint32, values []float32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttrib1fv(index, values)
}

func (R *RenderingContext) VertexAttrib2fv(index uint32, values []float32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttrib2fv(index, values)
}

func (R *RenderingContext) VertexAttrib3fv(index uint32, values []float32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttrib3fv(index, values)
}

func (R *RenderingContext) VertexAttrib4fv(index uint32, values []float32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttrib4fv(index, values)
}

func (R *RenderingContext) VertexAttribI4i(index uint32, x, y, z, w int32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttribI4i(index, x, y, z, w)
}

func (R *RenderingContext) VertexAttribI4ui(index, x, y, z, w uint32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttribI4ui(index, x, y, z, w)
}

func (R *RenderingContext) VertexAttribI4iv(index uint32, values []int32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttribI4iv(index, values)
}

func (R *RenderingContext) VertexAttribI4uiv(index uint32, values []uint32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttribI4uiv(index, values)
}

func (R *RenderingContext) VertexAttribPointer(index uint32, size int32, typ glone.Enum, normalized bool, stride, offset int32) {
	R.requiresVertexArray()
	R.requiresArrayBuffer()
	R.RenderingContext.VertexAttribPointer(index, size, typ, normalized, stride, offset)
}

func (R *RenderingContext) VertexAttribIPointer(index uint32, size int32, typ glone.Enum, stride, offset int32) {
	R.requiresVertexArray()
	R.requiresArrayBuffer()
	R.RenderingContext.VertexAttribIPointer(index, size, typ, stride, offset)
}

func (R *RenderingContext) VertexAttribDivisor(index, divisor uint32) {
	R.requiresVertexArray()
	R.RenderingContext.VertexAttribDivisor(index, divisor)
}

func (R *RenderingContext) GetVertexAttrib(index uint32, pname glone.Enum) any {
	R.requiresVertexArray()
	return R.RenderingContext.GetVertexAttrib(index, pname)
}

func (R *RenderingContext) GetVertexAttribOffset(index uint32, pname glone.Enum) int32 {
	R.requiresVertexArray()
	return R.RenderingContext.GetVertexAttribOffset(index, pname)
}

func (R *RenderingContext) BufferDataSize(target glone.Enum, size int32, usage glone.Enum) {
	R.requiresBuffer(target)
	R.RenderingContext.BufferDataSize(target, size, usage)
}

func (R *RenderingContext) BufferDataSrc(target glone.Enum, data glone.BufferSource, usage glone.Enum) {
	R.requiresBuffer(target)
	R.RenderingContext.BufferDataSrc(target, data, usage)
}

func (R *RenderingContext) BufferDataPix(target glone.Enum, srcData []byte, usage glone.Enum) {
	R.requiresBuffer(target)
	R.RenderingContext.BufferDataPix(target, srcData, usage)
}

func (R *RenderingContext) BufferSubDataSrc(target glone.Enum, offset int32, data glone.BufferSource) {
	R.requiresBuffer(target)
	R.RenderingContext.BufferSubDataSrc(target, offset, data)
}

func (R *RenderingContext) BufferSubDataPix(target glone.Enum, dstByteOffset int32, srcData []byte) {
	R.requiresBuffer(target)
	R.RenderingContext.BufferSubDataPix(target, dstByteOffset, srcData)
}

func (R *RenderingContext) GetBufferParameter(target, pname glone.Enum) any {
	R.requiresBuffer(target)
	return R.RenderingContext.GetBufferParameter(target, pname)
}

func (R *RenderingContext) CopyBufferSubData(readTarget, writeTarget glone.Enum, readOffset, writeOffset, size int32) {
	R.requiresBuffer(readTarget)
	R.requiresBuffer(writeTarget)
	R.RenderingContext.CopyBufferSubData(readTarget, writeTarget, readOffset, writeOffset, size)
}

func (R *RenderingContext) GetBufferSubData(target glone.Enum, srcByteOffset int32, dstBuffer []byte) {
	R.requiresBuffer(target)
	R.RenderingContext.GetBufferSubData(target, srcByteOffset, dstBuffer)
}

func (R *RenderingContext) DrawArrays(mode glone.Enum, first, count int32) {
	R.requiresVertexArray()
	R.requiresProgram()
	R.RenderingContext.DrawArrays(mode, first, count)
}

func (R *RenderingContext) DrawElements(mode glone.Enum, count int32, typ glone.Enum, offset int32) {
	R.requiresVertexArray()
	R.requiresElementArrayBuffer()
	R.requiresProgram()
	R.RenderingContext.DrawElements(mode, count, typ, offset)
}

func (R *RenderingContext) DrawArraysInstanced(mode glone.Enum, first, count, instanceCount int32) {
	R.requiresVertexArray()
	R.requiresProgram()
	R.RenderingContext.DrawArraysInstanced(mode, first, count, instanceCount)
}

func (R *RenderingContext) DrawElementsInstanced(mode glone.Enum, count int32, typ glone.Enum, offset, instanceCount int32) {
	R.requiresVertexArray()
	R.requiresElementArrayBuffer()
	R.requiresProgram()
	R.RenderingContext.DrawElementsInstanced(mode, count, typ, offset, instanceCount)
}

func (R *RenderingContext) DrawRangeElements(mode glone.Enum, start, end uint32, count int32, typ glone.Enum, offset int32) {
	R.requiresVertexArray()
	R.requiresElementArrayBuffer()
	R.requiresProgram()
	R.RenderingContext.DrawRangeElements(mode, start, end, count, typ, offset)
}
