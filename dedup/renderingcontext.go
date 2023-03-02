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

func (L *lazyUpdater[T]) setBoth(current T) {
	L.target = current
	L.current = current
}

type textureUnit struct {
	samplerBinding        lazyUpdater[glone.Sampler]
	texture2DBinding      lazyUpdater[glone.Texture]
	textureCubeMapBinding lazyUpdater[glone.Texture]
	texture3DBinding      lazyUpdater[glone.Texture]
	texture2DArrayBinding lazyUpdater[glone.Texture]
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
	framebufferBinding             lazyUpdater[glone.Framebuffer]
	drawFramebufferBinding         lazyUpdater[glone.Framebuffer]
	readFramebufferBinding         lazyUpdater[glone.Framebuffer]
	renderbufferBinding            lazyUpdater[glone.Renderbuffer]
	vertexArrayBinding             lazyUpdater[glone.VertexArray]
	programBinding                 lazyUpdater[glone.Program]
	transformFeedbackBinding       lazyUpdater[glone.TransformFeedback]
	activeTexture                  lazyUpdater[int]
	textureUnits                   []textureUnit
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

func (R *RenderingContext) requiresBufferTarget(target glone.Enum) {
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

func (R *RenderingContext) BindBufferBase(target glone.Enum, index uint32, buffer glone.Buffer) {
	switch target {
	case glone.TRANSFORM_FEEDBACK_BUFFER:
		R.requiresTransformFeedback()
		R.transformFeedbackBufferBinding.setBoth(buffer)
	case glone.UNIFORM_BUFFER:
		R.uniformBufferBinding.setBoth(buffer)
	default:
		return
	}
	R.RenderingContext.BindBufferBase(target, index, buffer)
}

func (R *RenderingContext) BindBufferRange(target glone.Enum, index uint32, buffer glone.Buffer, offset, size int32) {
	switch target {
	case glone.TRANSFORM_FEEDBACK_BUFFER:
		R.requiresTransformFeedback()
		R.transformFeedbackBufferBinding.setBoth(buffer)
	case glone.UNIFORM_BUFFER:
		R.uniformBufferBinding.setBoth(buffer)
	default:
		return
	}
	R.RenderingContext.BindBufferRange(target, index, buffer, offset, size)
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

func (R *RenderingContext) requiresFramebuffer() {
	if R.framebufferBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindFramebuffer(glone.FRAMEBUFFER, R.framebufferBinding.target)
	R.framebufferBinding.makeCurrent()
}

func (R *RenderingContext) requiresFramebufferTarget(target glone.Enum) {
	switch target {
	case glone.FRAMEBUFFER:
		R.requiresFramebuffer()
	case glone.READ_FRAMEBUFFER:
		R.requiresReadFramebuffer()
	case glone.DRAW_FRAMEBUFFER:
		R.requiresDrawFramebuffer()
	}
}

func (R *RenderingContext) BindFramebuffer(target glone.Enum, framebuffer glone.Framebuffer) {
	switch target {
	case glone.FRAMEBUFFER:
		R.framebufferBinding.target = framebuffer
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

func (R *RenderingContext) requiresRenderbufferTarget(target glone.Enum) {
	switch target {
	case glone.RENDERBUFFER:
		R.requiresRenderbuffer()
	}
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

func (R *RenderingContext) requiresTexture2D() {
	tu := R.getActiveTextureUnit()
	if tu.texture2DBinding.isCurrent() {
		return
	}
	R.requiresActiveTexture()
	R.RenderingContext.BindTexture(glone.TEXTURE_2D, tu.texture2DBinding.target)
	tu.texture2DBinding.makeCurrent()
}

func (R *RenderingContext) requiresTextureCubeMap() {
	tu := R.getActiveTextureUnit()
	if tu.textureCubeMapBinding.isCurrent() {
		return
	}
	R.requiresActiveTexture()
	R.RenderingContext.BindTexture(glone.TEXTURE_CUBE_MAP, tu.textureCubeMapBinding.target)
	tu.textureCubeMapBinding.makeCurrent()
}

func (R *RenderingContext) requiresTexture3D() {
	tu := R.getActiveTextureUnit()
	if tu.texture3DBinding.isCurrent() {
		return
	}
	R.requiresActiveTexture()
	R.RenderingContext.BindTexture(glone.TEXTURE_3D, tu.texture3DBinding.target)
	tu.texture3DBinding.makeCurrent()
}

func (R *RenderingContext) requiresTexture2DArray() {
	tu := R.getActiveTextureUnit()
	if tu.texture2DArrayBinding.isCurrent() {
		return
	}
	R.requiresActiveTexture()
	R.RenderingContext.BindTexture(glone.TEXTURE_2D_ARRAY, tu.texture2DArrayBinding.target)
	tu.texture2DArrayBinding.makeCurrent()
}

func (R *RenderingContext) requiresTextureTarget(target glone.Enum) {
	switch target {
	case glone.TEXTURE_2D:
		R.requiresTexture2D()
	case glone.TEXTURE_CUBE_MAP:
		R.requiresTextureCubeMap()
	case glone.TEXTURE_3D:
		R.requiresTexture3D()
	case glone.TEXTURE_2D_ARRAY:
		R.requiresTexture2DArray()
	}
}

func (R *RenderingContext) requiresTextures() {
	activeTexture := R.activeTexture.target
	for idx := range R.textureUnits {
		R.activeTexture.target = idx
		R.requiresSampler()
		R.requiresTexture2D()
		R.requiresTextureCubeMap()
		R.requiresTexture3D()
		R.requiresTexture2DArray()
	}
	R.activeTexture.target = activeTexture
}

func (R *RenderingContext) BindTexture(target glone.Enum, texture glone.Texture) {
	tu := R.getActiveTextureUnit()
	switch target {
	case glone.TEXTURE_2D:
		tu.texture2DBinding.target = texture
	case glone.TEXTURE_CUBE_MAP:
		tu.textureCubeMapBinding.target = texture
	case glone.TEXTURE_3D:
		tu.texture3DBinding.target = texture
	case glone.TEXTURE_2D_ARRAY:
		tu.texture2DArrayBinding.target = texture
	}
}

func (R *RenderingContext) BindVertexArray(array glone.VertexArray) {
	R.vertexArrayBinding.target = array
}

func (R *RenderingContext) requiresTransformFeedback() {
	if R.transformFeedbackBinding.isCurrent() {
		return
	}
	R.RenderingContext.BindTransformFeedback(glone.TRANSFORM_FEEDBACK, R.transformFeedbackBinding.target)
	R.transformFeedbackBinding.makeCurrent()
}

func (R *RenderingContext) requiresTransformFeedbackTarget(target glone.Enum) {
	switch target {
	case glone.TRANSFORM_FEEDBACK:
		R.requiresTransformFeedback()
	}
}

func (R *RenderingContext) BindTransformFeedback(target glone.Enum, tf glone.TransformFeedback) {
	switch target {
	case glone.TRANSFORM_FEEDBACK:
		R.transformFeedbackBinding.target = tf
	}
}

func (R *RenderingContext) getTextureUnit(unit int) *textureUnit {
	for unit >= len(R.textureUnits) {
		R.textureUnits = append(R.textureUnits, textureUnit{})
	}
	return &R.textureUnits[unit]
}

func (R *RenderingContext) getActiveTextureUnit() *textureUnit {
	return R.getTextureUnit(R.activeTexture.target)
}

func (R *RenderingContext) requiresSampler() {
	tu := R.getActiveTextureUnit()
	if tu.samplerBinding.isCurrent() {
		return
	}
	R.requiresActiveTexture()
	R.RenderingContext.BindSampler(uint32(R.activeTexture.target), tu.samplerBinding.target)
	tu.samplerBinding.makeCurrent()
}

func (R *RenderingContext) BindSampler(unit uint32, sampler glone.Sampler) {
	R.getTextureUnit(int(unit)).samplerBinding.target = sampler
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

func (R *RenderingContext) FramebufferRenderbuffer(target, attachment, renderbuffertarget glone.Enum, renderbuffer glone.Renderbuffer) {
	R.requiresFramebufferTarget(target)
	R.RenderingContext.FramebufferRenderbuffer(target, attachment, renderbuffertarget, renderbuffer)
}

func (R *RenderingContext) FramebufferTexture2D(target, attachment, textarget glone.Enum, texture glone.Texture, level int32) {
	R.requiresFramebufferTarget(target)
	R.RenderingContext.FramebufferTexture2D(target, attachment, textarget, texture, level)
}

func (R *RenderingContext) CheckFramebufferStatus(target glone.Enum) glone.Enum {
	R.requiresFramebufferTarget(target)
	return R.RenderingContext.CheckFramebufferStatus(target)
}

func (R *RenderingContext) GetFramebufferAttachmentParameter(target, attachment, pname glone.Enum) any {
	R.requiresFramebufferTarget(target)
	return R.RenderingContext.GetFramebufferAttachmentParameter(target, attachment, pname)
}

func (R *RenderingContext) FramebufferTextureLayer(target, attachment glone.Enum, texture glone.Texture, level, layer int32) {
	R.requiresFramebufferTarget(target)
	R.RenderingContext.FramebufferTextureLayer(target, attachment, texture, level, layer)
}

func (R *RenderingContext) InvalidateFramebuffer(target glone.Enum, attachments []glone.Enum) {
	R.requiresFramebufferTarget(target)
	R.RenderingContext.InvalidateFramebuffer(target, attachments)
}

func (R *RenderingContext) InvalidateSubFramebuffer(target glone.Enum, attachments []glone.Enum, x, y, width, height int32) {
	R.requiresFramebufferTarget(target)
	R.RenderingContext.InvalidateSubFramebuffer(target, attachments, x, y, width, height)
}

func (R *RenderingContext) BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int32, mask, filter glone.Enum) {
	R.requiresReadFramebuffer()
	R.requiresDrawFramebuffer()
	R.RenderingContext.BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter)
}

func (R *RenderingContext) RenderbufferStorage(target, internalformat glone.Enum, width, height int32) {
	R.requiresRenderbufferTarget(target)
	R.RenderingContext.RenderbufferStorage(target, internalformat, width, height)
}

func (R *RenderingContext) RenderbufferStorageMultisample(target glone.Enum, samples int32, internalformat glone.Enum, width, height int32) {
	R.requiresRenderbufferTarget(target)
	R.RenderingContext.RenderbufferStorageMultisample(target, samples, internalformat, width, height)
}

func (R *RenderingContext) GetRenderbufferParameter(target, pname glone.Enum) any {
	R.requiresRenderbufferTarget(target)
	return R.RenderingContext.GetRenderbufferParameter(target, pname)
}

func (R *RenderingContext) GetInternalFormatParameter(target, internalformat, pname glone.Enum) any {
	R.requiresRenderbufferTarget(target)
	return R.RenderingContext.GetInternalFormatParameter(target, internalformat, pname)
}

func (R *RenderingContext) BeginTransformFeedback(primitiveMode glone.Enum) {
	R.requiresTransformFeedback()
	R.RenderingContext.BeginTransformFeedback(primitiveMode)
}

func (R *RenderingContext) EndTransformFeedback() {
	R.requiresTransformFeedback()
	R.RenderingContext.EndTransformFeedback()
}

func (R *RenderingContext) PauseTransformFeedback() {
	R.requiresTransformFeedback()
	R.RenderingContext.PauseTransformFeedback()
}

func (R *RenderingContext) ResumeTransformFeedback() {
	R.requiresTransformFeedback()
	R.RenderingContext.ResumeTransformFeedback()
}

func (R *RenderingContext) Clear(mask glone.Enum) {
	R.requiresFramebuffer()
	R.requiresDrawFramebuffer()
	R.RenderingContext.Clear(mask)
}

func (R *RenderingContext) ClearBufferfv(buffer glone.Enum, drawbuffer int32, values []float32) {
	R.requiresFramebuffer()
	R.requiresDrawFramebuffer()
	R.RenderingContext.ClearBufferfv(buffer, drawbuffer, values)
}

func (R *RenderingContext) ClearBufferiv(buffer glone.Enum, drawbuffer int32, values []int32) {
	R.requiresFramebuffer()
	R.requiresDrawFramebuffer()
	R.RenderingContext.ClearBufferiv(buffer, drawbuffer, values)
}

func (R *RenderingContext) ClearBufferuiv(buffer glone.Enum, drawbuffer int32, values []uint32) {
	R.requiresFramebuffer()
	R.requiresDrawFramebuffer()
	R.RenderingContext.ClearBufferuiv(buffer, drawbuffer, values)
}

func (R *RenderingContext) ClearBufferfi(buffer glone.Enum, drawbuffer int32, depth float32, stencil int32) {
	R.requiresFramebuffer()
	R.requiresDrawFramebuffer()
	R.RenderingContext.ClearBufferfi(buffer, drawbuffer, depth, stencil)
}

func (R *RenderingContext) ReadPixelsPbo(x, y, width, height int32, format, typ glone.Enum, offset int32) {
	R.requiresFramebuffer()
	R.requiresReadFramebuffer()
	R.requiresBufferTarget(glone.PIXEL_PACK_BUFFER)
	R.RenderingContext.ReadPixelsPbo(x, y, width, height, format, typ, offset)
}

func (R *RenderingContext) ReadPixelsPix(x, y, width, height int32, format, typ glone.Enum, dstData []byte) {
	R.requiresFramebuffer()
	R.requiresReadFramebuffer()
	R.RenderingContext.ReadPixelsPix(x, y, width, height, format, typ, dstData)
}

func (R *RenderingContext) BufferDataSize(target glone.Enum, size int32, usage glone.Enum) {
	R.requiresBufferTarget(target)
	R.RenderingContext.BufferDataSize(target, size, usage)
}

func (R *RenderingContext) BufferDataSrc(target glone.Enum, data glone.BufferSource, usage glone.Enum) {
	R.requiresBufferTarget(target)
	R.RenderingContext.BufferDataSrc(target, data, usage)
}

func (R *RenderingContext) BufferDataPix(target glone.Enum, srcData []byte, usage glone.Enum) {
	R.requiresBufferTarget(target)
	R.RenderingContext.BufferDataPix(target, srcData, usage)
}

func (R *RenderingContext) BufferSubDataSrc(target glone.Enum, offset int32, data glone.BufferSource) {
	R.requiresBufferTarget(target)
	R.RenderingContext.BufferSubDataSrc(target, offset, data)
}

func (R *RenderingContext) BufferSubDataPix(target glone.Enum, dstByteOffset int32, srcData []byte) {
	R.requiresBufferTarget(target)
	R.RenderingContext.BufferSubDataPix(target, dstByteOffset, srcData)
}

func (R *RenderingContext) GetBufferParameter(target, pname glone.Enum) any {
	R.requiresBufferTarget(target)
	return R.RenderingContext.GetBufferParameter(target, pname)
}

func (R *RenderingContext) CopyBufferSubData(readTarget, writeTarget glone.Enum, readOffset, writeOffset, size int32) {
	R.requiresBufferTarget(readTarget)
	R.requiresBufferTarget(writeTarget)
	R.RenderingContext.CopyBufferSubData(readTarget, writeTarget, readOffset, writeOffset, size)
}

func (R *RenderingContext) GetBufferSubData(target glone.Enum, srcByteOffset int32, dstBuffer []byte) {
	R.requiresBufferTarget(target)
	R.RenderingContext.GetBufferSubData(target, srcByteOffset, dstBuffer)
}

func (R *RenderingContext) requiresActiveTexture() {
	if R.activeTexture.isCurrent() {
		return
	}
	R.RenderingContext.ActiveTexture(glone.Enum(R.activeTexture.target) + glone.TEXTURE0)
	R.activeTexture.makeCurrent()
}

func (R *RenderingContext) ActiveTexture(texture glone.Enum) {
	R.activeTexture.target = int(texture - glone.TEXTURE0)
}

func (R *RenderingContext) TexParameterf(target, pname glone.Enum, param float32) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexParameterf(target, pname, param)
}

func (R *RenderingContext) TexParameteri(target, pname glone.Enum, param int32) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexParameteri(target, pname, param)
}

func (R *RenderingContext) GetTexParameter(target, pname glone.Enum) any {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	return R.RenderingContext.GetTexParameter(target, pname)
}

func (R *RenderingContext) GenerateMipmap(target glone.Enum) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.GenerateMipmap(target)
}

func (R *RenderingContext) CopyTexImage2D(target glone.Enum, level int32, internalformat glone.Enum, x, y, width, height, border int32) {
	R.requiresFramebuffer()
	R.requiresReadFramebuffer()
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CopyTexImage2D(target, level, internalformat, x, y, width, height, border)
}

func (R *RenderingContext) CopyTexSubImage2D(target glone.Enum, level int32, xoffset, yoffset, x, y, width, height int32) {
	R.requiresFramebuffer()
	R.requiresReadFramebuffer()
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height)
}

func (R *RenderingContext) CopyTexSubImage3D(target glone.Enum, level, xoffset, yoffset, zoffset, x, y, width, height int32) {
	R.requiresFramebuffer()
	R.requiresReadFramebuffer()
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height)
}

func (R *RenderingContext) TexStorage2D(target glone.Enum, levels int32, internalformat glone.Enum, width, height int32) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexStorage2D(target, levels, internalformat, width, height)
}

func (R *RenderingContext) TexStorage3D(target glone.Enum, levels int32, internalformat glone.Enum, width, height, depth int32) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexStorage3D(target, levels, internalformat, width, height, depth)
}

func (R *RenderingContext) TexImage2DPbo(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, pboOffset int32) {
	R.requiresBufferTarget(glone.PIXEL_UNPACK_BUFFER)
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexImage2DPbo(target, level, internalformat, width, height, border, format, typ, pboOffset)
}

func (R *RenderingContext) TexImage2DSrc(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, source glone.TexImageSource) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexImage2DSrc(target, level, internalformat, width, height, border, format, typ, source)
}

func (R *RenderingContext) TexImage2DPix(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, srcData []byte) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexImage2DPix(target, level, internalformat, width, height, border, format, typ, srcData)
}

func (R *RenderingContext) TexSubImage2DPbo(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, pboOffset int32) {
	R.requiresBufferTarget(glone.PIXEL_UNPACK_BUFFER)
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexSubImage2DPbo(target, level, xoffset, yoffset, width, height, format, typ, pboOffset)
}

func (R *RenderingContext) TexSubImage2DSrc(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, source glone.TexImageSource) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexSubImage2DSrc(target, level, xoffset, yoffset, width, height, format, typ, source)
}

func (R *RenderingContext) TexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, srcData []byte) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexSubImage2DPix(target, level, xoffset, yoffset, width, height, format, typ, srcData)
}

func (R *RenderingContext) TexImage3DPbo(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, pboOffset int32) {
	R.requiresBufferTarget(glone.PIXEL_UNPACK_BUFFER)
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexImage3DPbo(target, level, internalformat, width, height, depth, border, format, typ, pboOffset)
}

func (R *RenderingContext) TexImage3DSrc(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, source glone.TexImageSource) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexImage3DSrc(target, level, internalformat, width, height, depth, border, format, typ, source)
}

func (R *RenderingContext) TexImage3DPix(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, srcData []byte) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexImage3DPix(target, level, internalformat, width, height, depth, border, format, typ, srcData)
}

func (R *RenderingContext) TexSubImage3DPbo(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, pboOffset int32) {
	R.requiresBufferTarget(glone.PIXEL_UNPACK_BUFFER)
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexSubImage3DPbo(target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, pboOffset)
}

func (R *RenderingContext) TexSubImage3DSrc(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, source glone.TexImageSource) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexSubImage3DSrc(target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, source)
}

func (R *RenderingContext) TexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, srcData []byte) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.TexSubImage3DPix(target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, srcData)
}

func (R *RenderingContext) CompressedTexImage2DSize(target glone.Enum, level int32, internalformat glone.Enum, width, height, border, imageSize, offset int32) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CompressedTexImage2DSize(target, level, internalformat, width, height, border, imageSize, offset)
}

func (R *RenderingContext) CompressedTexImage2DPix(target glone.Enum, level int32, internalformat glone.Enum, width, height, border int32, srcData []byte) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CompressedTexImage2DPix(target, level, internalformat, width, height, border, srcData)
}

func (R *RenderingContext) CompressedTexSubImage2DSize(target glone.Enum, level, xoffset, yoffset, width, height int32, format glone.Enum, imageSize, offset int32) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CompressedTexSubImage2DSize(target, level, xoffset, yoffset, width, height, format, imageSize, offset)
}

func (R *RenderingContext) CompressedTexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int32, format glone.Enum, srcData []byte) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CompressedTexSubImage2DPix(target, level, xoffset, yoffset, width, height, format, srcData)
}

func (R *RenderingContext) CompressedTexImage3DSize(target glone.Enum, level int32, internalformat glone.Enum, width, height, depth, border, imageSize, offset int32) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CompressedTexImage3DSize(target, level, internalformat, width, height, depth, border, imageSize, offset)
}

func (R *RenderingContext) CompressedTexImage3DPix(target glone.Enum, level int32, internalformat glone.Enum, width, height, depth, border int32, srcData []byte) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CompressedTexImage3DPix(target, level, internalformat, width, height, depth, border, srcData)
}

func (R *RenderingContext) CompressedTexSubImage3DSize(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format glone.Enum, imageSize, offset int32) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CompressedTexSubImage3DSize(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, offset)
}

func (R *RenderingContext) CompressedTexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format glone.Enum, srcData []byte) {
	R.requiresActiveTexture()
	R.requiresTextureTarget(target)
	R.RenderingContext.CompressedTexSubImage3DPix(target, level, xoffset, yoffset, zoffset, width, height, depth, format, srcData)
}

func (R *RenderingContext) requireForDraw() {
	R.requiresFramebuffer()
	R.requiresDrawFramebuffer()
	R.requiresVertexArray()
	R.requiresProgram()
	R.requiresTextures()
}

func (R *RenderingContext) DrawArrays(mode glone.Enum, first, count int32) {
	R.requireForDraw()
	R.RenderingContext.DrawArrays(mode, first, count)
}

func (R *RenderingContext) DrawElements(mode glone.Enum, count int32, typ glone.Enum, offset int32) {
	R.requireForDraw()
	R.requiresElementArrayBuffer()
	R.RenderingContext.DrawElements(mode, count, typ, offset)
}

func (R *RenderingContext) DrawArraysInstanced(mode glone.Enum, first, count, instanceCount int32) {
	R.requireForDraw()
	R.RenderingContext.DrawArraysInstanced(mode, first, count, instanceCount)
}

func (R *RenderingContext) DrawElementsInstanced(mode glone.Enum, count int32, typ glone.Enum, offset, instanceCount int32) {
	R.requireForDraw()
	R.requiresElementArrayBuffer()
	R.RenderingContext.DrawElementsInstanced(mode, count, typ, offset, instanceCount)
}

func (R *RenderingContext) DrawRangeElements(mode glone.Enum, start, end uint32, count int32, typ glone.Enum, offset int32) {
	R.requireForDraw()
	R.requiresElementArrayBuffer()
	R.RenderingContext.DrawRangeElements(mode, start, end, count, typ, offset)
}
