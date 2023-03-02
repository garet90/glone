package debugger

import (
	"fmt"
	"github.com/garet90/glone"
)

type RenderingContext struct {
	glone.RenderingContext
}

func MakeRenderingContext(context glone.RenderingContext) RenderingContext {
	return RenderingContext{
		context,
	}
}

func NewRenderingContext(context glone.RenderingContext) *RenderingContext {
	return &RenderingContext{
		context,
	}
}

func (R RenderingContext) checkError() {
	err := R.RenderingContext.GetError()
	if err != glone.NO_ERROR {
		panic(fmt.Sprintf("gl error: %d", err))
	}
}

func (R RenderingContext) CreateBuffer() glone.Buffer {
	defer R.checkError()
	return R.RenderingContext.CreateBuffer()
}

func (R RenderingContext) CreateFramebuffer() glone.Framebuffer {
	defer R.checkError()
	return R.RenderingContext.CreateFramebuffer()
}

func (R RenderingContext) CreateProgram() glone.Program {
	defer R.checkError()
	return R.RenderingContext.CreateProgram()
}

func (R RenderingContext) CreateRenderbuffer() glone.Renderbuffer {
	defer R.checkError()
	return R.RenderingContext.CreateRenderbuffer()
}

func (R RenderingContext) CreateShader(typ glone.Enum) glone.Shader {
	defer R.checkError()
	return R.RenderingContext.CreateShader(typ)
}

func (R RenderingContext) CreateTexture() glone.Texture {
	defer R.checkError()
	return R.RenderingContext.CreateTexture()
}

func (R RenderingContext) CreateVertexArray() glone.VertexArray {
	defer R.checkError()
	return R.RenderingContext.CreateVertexArray()
}

func (R RenderingContext) CreateTransformFeedback() glone.TransformFeedback {
	defer R.checkError()
	return R.RenderingContext.CreateTransformFeedback()
}

func (R RenderingContext) CreateSampler() glone.Sampler {
	defer R.checkError()
	return R.RenderingContext.CreateSampler()
}

func (R RenderingContext) CreateQuery() glone.Query {
	defer R.checkError()
	return R.RenderingContext.CreateQuery()
}

func (R RenderingContext) BindAttribLocation(program glone.Program, index uint32, name string) {
	defer R.checkError()
	R.RenderingContext.BindAttribLocation(program, index, name)
}

func (R RenderingContext) BindBuffer(target glone.Enum, buffer glone.Buffer) {
	defer R.checkError()
	R.RenderingContext.BindBuffer(target, buffer)
}

func (R RenderingContext) BindBufferBase(target glone.Enum, index uint32, buffer glone.Buffer) {
	defer R.checkError()
	R.RenderingContext.BindBufferBase(target, index, buffer)
}

func (R RenderingContext) BindBufferRange(target glone.Enum, index uint32, buffer glone.Buffer, offset, size int32) {
	defer R.checkError()
	R.RenderingContext.BindBufferRange(target, index, buffer, offset, size)
}

func (R RenderingContext) BindFramebuffer(target glone.Enum, framebuffer glone.Framebuffer) {
	defer R.checkError()
	R.RenderingContext.BindFramebuffer(target, framebuffer)
}

func (R RenderingContext) BindRenderbuffer(target glone.Enum, renderbuffer glone.Renderbuffer) {
	defer R.checkError()
	R.RenderingContext.BindRenderbuffer(target, renderbuffer)
}

func (R RenderingContext) BindTexture(target glone.Enum, texture glone.Texture) {
	defer R.checkError()
	R.RenderingContext.BindTexture(target, texture)
}

func (R RenderingContext) BindVertexArray(array glone.VertexArray) {
	defer R.checkError()
	R.RenderingContext.BindVertexArray(array)
}

func (R RenderingContext) BindTransformFeedback(target glone.Enum, tf glone.TransformFeedback) {
	defer R.checkError()
	R.RenderingContext.BindTransformFeedback(target, tf)
}

func (R RenderingContext) BindSampler(unit uint32, sampler glone.Sampler) {
	defer R.checkError()
	R.RenderingContext.BindSampler(unit, sampler)
}

func (R RenderingContext) IsBuffer(buffer glone.Buffer) bool {
	defer R.checkError()
	return R.RenderingContext.IsBuffer(buffer)
}

func (R RenderingContext) IsEnabled(cap glone.Enum) bool {
	defer R.checkError()
	return R.RenderingContext.IsEnabled(cap)
}

func (R RenderingContext) IsFramebuffer(framebuffer glone.Framebuffer) bool {
	defer R.checkError()
	return R.RenderingContext.IsFramebuffer(framebuffer)
}

func (R RenderingContext) IsProgram(program glone.Program) bool {
	defer R.checkError()
	return R.RenderingContext.IsProgram(program)
}

func (R RenderingContext) IsRenderbuffer(renderbuffer glone.Renderbuffer) bool {
	defer R.checkError()
	return R.RenderingContext.IsRenderbuffer(renderbuffer)
}

func (R RenderingContext) IsShader(shader glone.Shader) bool {
	defer R.checkError()
	return R.RenderingContext.IsShader(shader)
}

func (R RenderingContext) IsTexture(texture glone.Texture) bool {
	defer R.checkError()
	return R.RenderingContext.IsTexture(texture)
}

func (R RenderingContext) IsVertexArray(vertexArray glone.VertexArray) bool {
	defer R.checkError()
	return R.RenderingContext.IsVertexArray(vertexArray)
}

func (R RenderingContext) IsTransformFeedback(tf glone.TransformFeedback) bool {
	defer R.checkError()
	return R.RenderingContext.IsTransformFeedback(tf)
}

func (R RenderingContext) IsSampler(sampler glone.Sampler) bool {
	defer R.checkError()
	return R.RenderingContext.IsSampler(sampler)
}

func (R RenderingContext) IsQuery(query glone.Query) bool {
	defer R.checkError()
	return R.RenderingContext.IsQuery(query)
}

func (R RenderingContext) IsSync(sync glone.Sync) bool {
	defer R.checkError()
	return R.RenderingContext.IsSync(sync)
}

func (R RenderingContext) DeleteBuffer(buffer glone.Buffer) {
	defer R.checkError()
	R.RenderingContext.DeleteBuffer(buffer)
}

func (R RenderingContext) DeleteFramebuffer(framebuffer glone.Framebuffer) {
	defer R.checkError()
	R.RenderingContext.DeleteFramebuffer(framebuffer)
}

func (R RenderingContext) DeleteProgram(program glone.Program) {
	defer R.checkError()
	R.RenderingContext.DeleteProgram(program)
}

func (R RenderingContext) DeleteRenderbuffer(renderbuffer glone.Renderbuffer) {
	defer R.checkError()
	R.RenderingContext.DeleteRenderbuffer(renderbuffer)
}

func (R RenderingContext) DeleteShader(shader glone.Shader) {
	defer R.checkError()
	R.RenderingContext.DeleteShader(shader)
}

func (R RenderingContext) DeleteTexture(texture glone.Texture) {
	defer R.checkError()
	R.RenderingContext.DeleteTexture(texture)
}

func (R RenderingContext) DeleteVertexArray(vertexArray glone.VertexArray) {
	defer R.checkError()
	R.RenderingContext.DeleteVertexArray(vertexArray)
}

func (R RenderingContext) DeleteTransformFeedback(tf glone.TransformFeedback) {
	defer R.checkError()
	R.RenderingContext.DeleteTransformFeedback(tf)
}

func (R RenderingContext) DeleteSampler(sampler glone.Sampler) {
	defer R.checkError()
	R.RenderingContext.DeleteSampler(sampler)
}

func (R RenderingContext) DeleteQuery(query glone.Query) {
	defer R.checkError()
	R.RenderingContext.DeleteQuery(query)
}

func (R RenderingContext) DeleteSync(sync glone.Sync) {
	defer R.checkError()
	R.RenderingContext.DeleteSync(sync)
}

func (R RenderingContext) ShaderSource(shader glone.Shader, source string) {
	defer R.checkError()
	R.RenderingContext.ShaderSource(shader, source)
}

func (R RenderingContext) CompileShader(shader glone.Shader) {
	defer R.checkError()
	R.RenderingContext.CompileShader(shader)
}

func (R RenderingContext) GetShaderSource(shader glone.Shader) string {
	defer R.checkError()
	return R.RenderingContext.GetShaderSource(shader)
}

func (R RenderingContext) GetShaderParameter(shader glone.Shader, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetShaderParameter(shader, pname)
}

func (R RenderingContext) GetShaderPrecisionFormat(shadertype, precisiontype glone.Enum) glone.ShaderPrecisionFormat {
	defer R.checkError()
	return R.RenderingContext.GetShaderPrecisionFormat(shadertype, precisiontype)
}

func (R RenderingContext) GetShaderInfoLog(shader glone.Shader) string {
	defer R.checkError()
	return R.RenderingContext.GetShaderInfoLog(shader)
}

func (R RenderingContext) AttachShader(program glone.Program, shader glone.Shader) {
	defer R.checkError()
	R.RenderingContext.AttachShader(program, shader)
}

func (R RenderingContext) DetachShader(program glone.Program, shader glone.Shader) {
	defer R.checkError()
	R.RenderingContext.DetachShader(program, shader)
}

func (R RenderingContext) LinkProgram(program glone.Program) {
	defer R.checkError()
	R.RenderingContext.LinkProgram(program)
}

func (R RenderingContext) ValidateProgram(program glone.Program) {
	defer R.checkError()
	R.RenderingContext.ValidateProgram(program)
}

func (R RenderingContext) UseProgram(program glone.Program) {
	defer R.checkError()
	R.RenderingContext.UseProgram(program)
}

func (R RenderingContext) GetActiveAttrib(program glone.Program, index uint32) glone.ActiveInfo {
	defer R.checkError()
	return R.RenderingContext.GetActiveAttrib(program, index)
}

func (R RenderingContext) GetActiveUniform(program glone.Program, index uint32) glone.ActiveInfo {
	defer R.checkError()
	return R.RenderingContext.GetActiveUniform(program, index)
}

func (R RenderingContext) GetAttachedShaders(program glone.Program) []glone.Shader {
	defer R.checkError()
	return R.RenderingContext.GetAttachedShaders(program)
}

func (R RenderingContext) GetAttribLocation(program glone.Program, name string) int32 {
	defer R.checkError()
	return R.RenderingContext.GetAttribLocation(program, name)
}

func (R RenderingContext) GetProgramInfoLog(program glone.Program) string {
	defer R.checkError()
	return R.RenderingContext.GetProgramInfoLog(program)
}

func (R RenderingContext) GetUniform(program glone.Program, location glone.UniformLocation) any {
	defer R.checkError()
	return R.RenderingContext.GetUniform(program, location)
}

func (R RenderingContext) GetUniformLocation(program glone.Program, name string) glone.UniformLocation {
	defer R.checkError()
	return R.RenderingContext.GetUniformLocation(program, name)
}

func (R RenderingContext) GetFragDataLocation(program glone.Program, name string) int32 {
	defer R.checkError()
	return R.RenderingContext.GetFragDataLocation(program, name)
}

func (R RenderingContext) GetProgramParameter(program glone.Program, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetProgramParameter(program, pname)
}

func (R RenderingContext) TransformFeedbackVaryings(program glone.Program, varyings []string, bufferMode glone.Enum) {
	defer R.checkError()
	R.RenderingContext.TransformFeedbackVaryings(program, varyings, bufferMode)
}

func (R RenderingContext) GetTransformFeedbackVarying(program glone.Program, index uint32) glone.ActiveInfo {
	defer R.checkError()
	return R.RenderingContext.GetTransformFeedbackVarying(program, index)
}

func (R RenderingContext) Uniform1f(location glone.UniformLocation, x float32) {
	defer R.checkError()
	R.RenderingContext.Uniform1f(location, x)
}

func (R RenderingContext) Uniform2f(location glone.UniformLocation, x, y float32) {
	defer R.checkError()
	R.RenderingContext.Uniform2f(location, x, y)
}

func (R RenderingContext) Uniform3f(location glone.UniformLocation, x, y, z float32) {
	defer R.checkError()
	R.RenderingContext.Uniform3f(location, x, y, z)
}

func (R RenderingContext) Uniform4f(location glone.UniformLocation, x, y, z, w float32) {
	defer R.checkError()
	R.RenderingContext.Uniform4f(location, x, y, z, w)
}

func (R RenderingContext) Uniform1i(location glone.UniformLocation, x int32) {
	defer R.checkError()
	R.RenderingContext.Uniform1i(location, x)
}

func (R RenderingContext) Uniform2i(location glone.UniformLocation, x, y int32) {
	defer R.checkError()
	R.RenderingContext.Uniform2i(location, x, y)
}

func (R RenderingContext) Uniform3i(location glone.UniformLocation, x, y, z int32) {
	defer R.checkError()
	R.RenderingContext.Uniform3i(location, x, y, z)
}

func (R RenderingContext) Uniform4i(location glone.UniformLocation, x, y, z, w int32) {
	defer R.checkError()
	R.RenderingContext.Uniform4i(location, x, y, z, w)
}

func (R RenderingContext) Uniform1ui(location glone.UniformLocation, v0 uint32) {
	defer R.checkError()
	R.RenderingContext.Uniform1ui(location, v0)
}

func (R RenderingContext) Uniform2ui(location glone.UniformLocation, v0, v1 uint32) {
	defer R.checkError()
	R.RenderingContext.Uniform2ui(location, v0, v1)
}

func (R RenderingContext) Uniform3ui(location glone.UniformLocation, v0, v1, v2 uint32) {
	defer R.checkError()
	R.RenderingContext.Uniform3ui(location, v0, v1, v2)
}

func (R RenderingContext) Uniform4ui(location glone.UniformLocation, v0, v1, v2, v3 uint32) {
	defer R.checkError()
	R.RenderingContext.Uniform4ui(location, v0, v1, v2, v3)
}

func (R RenderingContext) Uniform1fv(location glone.UniformLocation, v []float32) {
	defer R.checkError()
	R.RenderingContext.Uniform1fv(location, v)
}

func (R RenderingContext) Uniform2fv(location glone.UniformLocation, v []float32) {
	defer R.checkError()
	R.RenderingContext.Uniform2fv(location, v)
}

func (R RenderingContext) Uniform3fv(location glone.UniformLocation, v []float32) {
	defer R.checkError()
	R.RenderingContext.Uniform3fv(location, v)
}

func (R RenderingContext) Uniform4fv(location glone.UniformLocation, v []float32) {
	defer R.checkError()
	R.RenderingContext.Uniform4fv(location, v)
}

func (R RenderingContext) Uniform1iv(location glone.UniformLocation, v []int32) {
	defer R.checkError()
	R.RenderingContext.Uniform1iv(location, v)
}

func (R RenderingContext) Uniform2iv(location glone.UniformLocation, v []int32) {
	defer R.checkError()
	R.RenderingContext.Uniform2iv(location, v)
}

func (R RenderingContext) Uniform3iv(location glone.UniformLocation, v []int32) {
	defer R.checkError()
	R.RenderingContext.Uniform3iv(location, v)
}

func (R RenderingContext) Uniform4iv(location glone.UniformLocation, v []int32) {
	defer R.checkError()
	R.RenderingContext.Uniform4iv(location, v)
}

func (R RenderingContext) Uniform1uiv(location glone.UniformLocation, data []uint32) {
	defer R.checkError()
	R.RenderingContext.Uniform1uiv(location, data)
}

func (R RenderingContext) Uniform2uiv(location glone.UniformLocation, data []uint32) {
	defer R.checkError()
	R.RenderingContext.Uniform2uiv(location, data)
}

func (R RenderingContext) Uniform3uiv(location glone.UniformLocation, data []uint32) {
	defer R.checkError()
	R.RenderingContext.Uniform3uiv(location, data)
}

func (R RenderingContext) Uniform4uiv(location glone.UniformLocation, data []uint32) {
	defer R.checkError()
	R.RenderingContext.Uniform4uiv(location, data)
}

func (R RenderingContext) UniformMatrix2fv(location glone.UniformLocation, transpose bool, value []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix2fv(location, transpose, value)
}

func (R RenderingContext) UniformMatrix3fv(location glone.UniformLocation, transpose bool, value []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix3fv(location, transpose, value)
}

func (R RenderingContext) UniformMatrix4fv(location glone.UniformLocation, transpose bool, value []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix4fv(location, transpose, value)
}

func (R RenderingContext) UniformMatrix3x2fv(location glone.UniformLocation, transpose bool, data []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix3x2fv(location, transpose, data)
}

func (R RenderingContext) UniformMatrix4x2fv(location glone.UniformLocation, transpose bool, data []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix4x2fv(location, transpose, data)
}

func (R RenderingContext) UniformMatrix2x3fv(location glone.UniformLocation, transpose bool, data []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix2x3fv(location, transpose, data)
}

func (R RenderingContext) UniformMatrix4x3fv(location glone.UniformLocation, transpose bool, data []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix4x3fv(location, transpose, data)
}

func (R RenderingContext) UniformMatrix2x4fv(location glone.UniformLocation, transpose bool, data []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix2x4fv(location, transpose, data)
}

func (R RenderingContext) UniformMatrix3x4fv(location glone.UniformLocation, transpose bool, data []float32) {
	defer R.checkError()
	R.RenderingContext.UniformMatrix3x4fv(location, transpose, data)
}

func (R RenderingContext) EnableVertexAttribArray(index uint32) {
	defer R.checkError()
	R.RenderingContext.EnableVertexAttribArray(index)
}

func (R RenderingContext) DisableVertexAttribArray(index uint32) {
	defer R.checkError()
	R.RenderingContext.DisableVertexAttribArray(index)
}

func (R RenderingContext) VertexAttrib1f(index uint32, x float32) {
	defer R.checkError()
	R.RenderingContext.VertexAttrib1f(index, x)
}

func (R RenderingContext) VertexAttrib2f(index uint32, x, y float32) {
	defer R.checkError()
	R.RenderingContext.VertexAttrib2f(index, x, y)
}

func (R RenderingContext) VertexAttrib3f(index uint32, x, y, z float32) {
	defer R.checkError()
	R.RenderingContext.VertexAttrib3f(index, x, y, z)
}

func (R RenderingContext) VertexAttrib4f(index uint32, x, y, z, w float32) {
	defer R.checkError()
	R.RenderingContext.VertexAttrib4f(index, x, y, z, w)
}

func (R RenderingContext) VertexAttrib1fv(index uint32, values []float32) {
	defer R.checkError()
	R.RenderingContext.VertexAttrib1fv(index, values)
}

func (R RenderingContext) VertexAttrib2fv(index uint32, values []float32) {
	defer R.checkError()
	R.RenderingContext.VertexAttrib2fv(index, values)
}

func (R RenderingContext) VertexAttrib3fv(index uint32, values []float32) {
	defer R.checkError()
	R.RenderingContext.VertexAttrib3fv(index, values)
}

func (R RenderingContext) VertexAttrib4fv(index uint32, values []float32) {
	defer R.checkError()
	R.RenderingContext.VertexAttrib4fv(index, values)
}

func (R RenderingContext) VertexAttribI4i(index uint32, x, y, z, w int32) {
	defer R.checkError()
	R.RenderingContext.VertexAttribI4i(index, x, y, z, w)
}

func (R RenderingContext) VertexAttribI4ui(index, x, y, z, w uint32) {
	defer R.checkError()
	R.RenderingContext.VertexAttribI4ui(index, x, y, z, w)
}

func (R RenderingContext) VertexAttribI4iv(index uint32, values []int32) {
	defer R.checkError()
	R.RenderingContext.VertexAttribI4iv(index, values)
}

func (R RenderingContext) VertexAttribI4uiv(index uint32, values []uint32) {
	defer R.checkError()
	R.RenderingContext.VertexAttribI4uiv(index, values)
}

func (R RenderingContext) VertexAttribPointer(index uint32, size int32, typ glone.Enum, normalized bool, stride, offset int32) {
	defer R.checkError()
	R.RenderingContext.VertexAttribPointer(index, size, typ, normalized, stride, offset)
}

func (R RenderingContext) VertexAttribIPointer(index uint32, size int32, typ glone.Enum, stride, offset int32) {
	defer R.checkError()
	R.RenderingContext.VertexAttribIPointer(index, size, typ, stride, offset)
}

func (R RenderingContext) VertexAttribDivisor(index, divisor uint32) {
	defer R.checkError()
	R.RenderingContext.VertexAttribDivisor(index, divisor)
}

func (R RenderingContext) GetVertexAttrib(index uint32, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetVertexAttrib(index, pname)
}

func (R RenderingContext) GetVertexAttribOffset(index uint32, pname glone.Enum) int32 {
	defer R.checkError()
	return R.RenderingContext.GetVertexAttribOffset(index, pname)
}

func (R RenderingContext) FramebufferRenderbuffer(target, attachment, renderbuffertarget glone.Enum, renderbuffer glone.Renderbuffer) {
	defer R.checkError()
	R.RenderingContext.FramebufferRenderbuffer(target, attachment, renderbuffertarget, renderbuffer)
}

func (R RenderingContext) FramebufferTexture2D(target, attachment, textarget glone.Enum, texture glone.Texture, level int32) {
	defer R.checkError()
	R.RenderingContext.FramebufferTexture2D(target, attachment, textarget, texture, level)
}

func (R RenderingContext) CheckFramebufferStatus(target glone.Enum) glone.Enum {
	defer R.checkError()
	return R.RenderingContext.CheckFramebufferStatus(target)
}

func (R RenderingContext) GetFramebufferAttachmentParameter(target, attachment, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetFramebufferAttachmentParameter(target, attachment, pname)
}

func (R RenderingContext) FramebufferTextureLayer(target, attachment glone.Enum, texture glone.Texture, level, layer int32) {
	defer R.checkError()
	R.RenderingContext.FramebufferTextureLayer(target, attachment, texture, level, layer)
}

func (R RenderingContext) InvalidateFramebuffer(target glone.Enum, attachments []glone.Enum) {
	defer R.checkError()
	R.RenderingContext.InvalidateFramebuffer(target, attachments)
}

func (R RenderingContext) InvalidateSubFramebuffer(target glone.Enum, attachments []glone.Enum, x, y, width, height int32) {
	defer R.checkError()
	R.RenderingContext.InvalidateSubFramebuffer(target, attachments, x, y, width, height)
}

func (R RenderingContext) BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int32, mask, filter glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter)
}

func (R RenderingContext) RenderbufferStorage(target, internalformat glone.Enum, width, height int32) {
	defer R.checkError()
	R.RenderingContext.RenderbufferStorage(target, internalformat, width, height)
}

func (R RenderingContext) RenderbufferStorageMultisample(target glone.Enum, samples int32, internalformat glone.Enum, width, height int32) {
	defer R.checkError()
	R.RenderingContext.RenderbufferStorageMultisample(target, samples, internalformat, width, height)
}

func (R RenderingContext) GetRenderbufferParameter(target, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetRenderbufferParameter(target, pname)
}

func (R RenderingContext) GetInternalFormatParameter(target, internalformat, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetInternalFormatParameter(target, internalformat, pname)
}

func (R RenderingContext) ReadPixelsPbo(x, y, width, height int32, format, typ glone.Enum, offset int32) {
	defer R.checkError()
	R.RenderingContext.ReadPixelsPbo(x, y, width, height, format, typ, offset)
}

func (R RenderingContext) ReadPixelsPix(x, y, width, height int32, format, typ glone.Enum, dstData []byte) {
	defer R.checkError()
	R.RenderingContext.ReadPixelsPix(x, y, width, height, format, typ, dstData)
}

func (R RenderingContext) BufferDataSize(target glone.Enum, size int32, usage glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BufferDataSize(target, size, usage)
}

func (R RenderingContext) BufferDataSrc(target glone.Enum, data glone.BufferSource, usage glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BufferDataSrc(target, data, usage)
}

func (R RenderingContext) BufferDataPix(target glone.Enum, srcData []byte, usage glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BufferDataPix(target, srcData, usage)
}

func (R RenderingContext) BufferSubDataSrc(target glone.Enum, offset int32, data glone.BufferSource) {
	defer R.checkError()
	R.RenderingContext.BufferSubDataSrc(target, offset, data)
}

func (R RenderingContext) BufferSubDataPix(target glone.Enum, dstByteOffset int32, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.BufferSubDataPix(target, dstByteOffset, srcData)
}

func (R RenderingContext) GetBufferParameter(target, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetBufferParameter(target, pname)
}

func (R RenderingContext) CopyBufferSubData(readTarget, writeTarget glone.Enum, readOffset, writeOffset, size int32) {
	defer R.checkError()
	R.RenderingContext.CopyBufferSubData(readTarget, writeTarget, readOffset, writeOffset, size)
}

func (R RenderingContext) GetBufferSubData(target glone.Enum, srcByteOffset int32, dstBuffer []byte) {
	defer R.checkError()
	R.RenderingContext.GetBufferSubData(target, srcByteOffset, dstBuffer)
}

func (R RenderingContext) ActiveTexture(texture glone.Enum) {
	defer R.checkError()
	R.RenderingContext.ActiveTexture(texture)
}

func (R RenderingContext) TexParameterf(target, pname glone.Enum, param float32) {
	defer R.checkError()
	R.RenderingContext.TexParameterf(target, pname, param)
}

func (R RenderingContext) TexParameteri(target, pname glone.Enum, param int32) {
	defer R.checkError()
	R.RenderingContext.TexParameteri(target, pname, param)
}

func (R RenderingContext) GetTexParameter(target, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetTexParameter(target, pname)
}

func (R RenderingContext) GenerateMipmap(target glone.Enum) {
	defer R.checkError()
	R.RenderingContext.GenerateMipmap(target)
}

func (R RenderingContext) CopyTexImage2D(target glone.Enum, level int32, internalformat glone.Enum, x, y, width, height, border int32) {
	defer R.checkError()
	R.RenderingContext.CopyTexImage2D(target, level, internalformat, x, y, width, height, border)
}

func (R RenderingContext) CopyTexSubImage2D(target glone.Enum, level int32, xoffset, yoffset, x, y, width, height int32) {
	defer R.checkError()
	R.RenderingContext.CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height)
}

func (R RenderingContext) CopyTexSubImage3D(target glone.Enum, level, xoffset, yoffset, zoffset, x, y, width, height int32) {
	defer R.checkError()
	R.RenderingContext.CopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height)
}

func (R RenderingContext) TexStorage2D(target glone.Enum, levels int32, internalformat glone.Enum, width, height int32) {
	defer R.checkError()
	R.RenderingContext.TexStorage2D(target, levels, internalformat, width, height)
}

func (R RenderingContext) TexStorage3D(target glone.Enum, levels int32, internalformat glone.Enum, width, height, depth int32) {
	defer R.checkError()
	R.RenderingContext.TexStorage3D(target, levels, internalformat, width, height, depth)
}

func (R RenderingContext) TexImage2DPbo(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, pboOffset int32) {
	defer R.checkError()
	R.RenderingContext.TexImage2DPbo(target, level, internalformat, width, height, border, format, typ, pboOffset)
}

func (R RenderingContext) TexImage2DSrc(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, source glone.TexImageSource) {
	defer R.checkError()
	R.RenderingContext.TexImage2DSrc(target, level, internalformat, width, height, border, format, typ, source)
}

func (R RenderingContext) TexImage2DPix(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.TexImage2DPix(target, level, internalformat, width, height, border, format, typ, srcData)
}

func (R RenderingContext) TexSubImage2DPbo(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, pboOffset int32) {
	defer R.checkError()
	R.RenderingContext.TexSubImage2DPbo(target, level, xoffset, yoffset, width, height, format, typ, pboOffset)
}

func (R RenderingContext) TexSubImage2DSrc(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, source glone.TexImageSource) {
	defer R.checkError()
	R.RenderingContext.TexSubImage2DSrc(target, level, xoffset, yoffset, width, height, format, typ, source)
}

func (R RenderingContext) TexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.TexSubImage2DPix(target, level, xoffset, yoffset, width, height, format, typ, srcData)
}

func (R RenderingContext) TexImage3DPbo(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, pboOffset int32) {
	defer R.checkError()
	R.RenderingContext.TexImage3DPbo(target, level, internalformat, width, height, depth, border, format, typ, pboOffset)
}

func (R RenderingContext) TexImage3DSrc(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, source glone.TexImageSource) {
	defer R.checkError()
	R.RenderingContext.TexImage3DSrc(target, level, internalformat, width, height, depth, border, format, typ, source)
}

func (R RenderingContext) TexImage3DPix(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.TexImage3DPix(target, level, internalformat, width, height, depth, border, format, typ, srcData)
}

func (R RenderingContext) TexSubImage3DPbo(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, pboOffset int32) {
	defer R.checkError()
	R.RenderingContext.TexSubImage3DPbo(target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, pboOffset)
}

func (R RenderingContext) TexSubImage3DSrc(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, source glone.TexImageSource) {
	defer R.checkError()
	R.RenderingContext.TexSubImage3DSrc(target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, source)
}

func (R RenderingContext) TexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.TexSubImage3DPix(target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, srcData)
}

func (R RenderingContext) CompressedTexImage2DSize(target glone.Enum, level int32, internalformat glone.Enum, width, height, border, imageSize, offset int32) {
	defer R.checkError()
	R.RenderingContext.CompressedTexImage2DSize(target, level, internalformat, width, height, border, imageSize, offset)
}

func (R RenderingContext) CompressedTexImage2DPix(target glone.Enum, level int32, internalformat glone.Enum, width, height, border int32, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.CompressedTexImage2DPix(target, level, internalformat, width, height, border, srcData)
}

func (R RenderingContext) CompressedTexSubImage2DSize(target glone.Enum, level, xoffset, yoffset, width, height int32, format glone.Enum, imageSize, offset int32) {
	defer R.checkError()
	R.RenderingContext.CompressedTexSubImage2DSize(target, level, xoffset, yoffset, width, height, format, imageSize, offset)
}

func (R RenderingContext) CompressedTexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int32, format glone.Enum, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.CompressedTexSubImage2DPix(target, level, xoffset, yoffset, width, height, format, srcData)
}

func (R RenderingContext) CompressedTexImage3DSize(target glone.Enum, level int32, internalformat glone.Enum, width, height, depth, border, imageSize, offset int32) {
	defer R.checkError()
	R.RenderingContext.CompressedTexImage3DSize(target, level, internalformat, width, height, depth, border, imageSize, offset)
}

func (R RenderingContext) CompressedTexImage3DPix(target glone.Enum, level int32, internalformat glone.Enum, width, height, depth, border int32, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.CompressedTexImage3DPix(target, level, internalformat, width, height, depth, border, srcData)
}

func (R RenderingContext) CompressedTexSubImage3DSize(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format glone.Enum, imageSize, offset int32) {
	defer R.checkError()
	R.RenderingContext.CompressedTexSubImage3DSize(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, offset)
}

func (R RenderingContext) CompressedTexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format glone.Enum, srcData []byte) {
	defer R.checkError()
	R.RenderingContext.CompressedTexSubImage3DPix(target, level, xoffset, yoffset, zoffset, width, height, depth, format, srcData)
}

func (R RenderingContext) BeginQuery(target glone.Enum, query glone.Query) {
	defer R.checkError()
	R.RenderingContext.BeginQuery(target, query)
}

func (R RenderingContext) EndQuery(target glone.Enum) {
	defer R.checkError()
	R.RenderingContext.EndQuery(target)
}

func (R RenderingContext) GetQuery(target, pname glone.Enum) glone.Query {
	defer R.checkError()
	return R.RenderingContext.GetQuery(target, pname)
}

func (R RenderingContext) GetQueryParameter(query glone.Query, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetQueryParameter(query, pname)
}

func (R RenderingContext) SamplerParameteri(sampler glone.Sampler, pname glone.Enum, param int32) {
	defer R.checkError()
	R.RenderingContext.SamplerParameteri(sampler, pname, param)
}

func (R RenderingContext) SamplerParameterf(sampler glone.Sampler, pname glone.Enum, param float32) {
	defer R.checkError()
	R.RenderingContext.SamplerParameterf(sampler, pname, param)
}

func (R RenderingContext) GetSamplerParameter(sampler glone.Sampler, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetSamplerParameter(sampler, pname)
}

func (R RenderingContext) FenceSync(condition, flags glone.Enum) glone.Sync {
	defer R.checkError()
	return R.RenderingContext.FenceSync(condition, flags)
}

func (R RenderingContext) ClientWaitSync(sync glone.Sync, flags glone.Enum, timeout uint64) {
	defer R.checkError()
	R.RenderingContext.ClientWaitSync(sync, flags, timeout)
}

func (R RenderingContext) WaitSync(sync glone.Sync, flags glone.Enum, timeout int64) {
	defer R.checkError()
	R.RenderingContext.WaitSync(sync, flags, timeout)
}

func (R RenderingContext) GetSyncParameter(sync glone.Sync, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetSyncParameter(sync, pname)
}

func (R RenderingContext) BeginTransformFeedback(primitiveMode glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BeginTransformFeedback(primitiveMode)
}

func (R RenderingContext) EndTransformFeedback() {
	defer R.checkError()
	R.RenderingContext.EndTransformFeedback()
}

func (R RenderingContext) PauseTransformFeedback() {
	defer R.checkError()
	R.RenderingContext.PauseTransformFeedback()
}

func (R RenderingContext) ResumeTransformFeedback() {
	defer R.checkError()
	R.RenderingContext.ResumeTransformFeedback()
}

func (R RenderingContext) GetIndexedParameter(target glone.Enum, index uint32) any {
	defer R.checkError()
	return R.RenderingContext.GetIndexedParameter(target, index)
}

func (R RenderingContext) GetUniformIndices(program glone.Program, uniformNames []string) []uint32 {
	defer R.checkError()
	return R.RenderingContext.GetUniformIndices(program, uniformNames)
}

func (R RenderingContext) GetActiveUniforms(program glone.Program, uniformIndices []uint32, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetActiveUniforms(program, uniformIndices, pname)
}

func (R RenderingContext) GetUniformBlockIndex(program glone.Program, uniformBlockName string) uint32 {
	defer R.checkError()
	return R.RenderingContext.GetUniformBlockIndex(program, uniformBlockName)
}

func (R RenderingContext) GetActiveUniformBlockParameter(program glone.Program, uniformBlockIndex uint32, pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetActiveUniformBlockParameter(program, uniformBlockIndex, pname)
}

func (R RenderingContext) GetActiveUniformBlockName(program glone.Program, uniformBlockIndex uint32) string {
	defer R.checkError()
	return R.RenderingContext.GetActiveUniformBlockName(program, uniformBlockIndex)
}

func (R RenderingContext) UniformBlockBinding(program glone.Program, uniformBlockIndex, uniformBlockBinding uint32) {
	defer R.checkError()
	R.RenderingContext.UniformBlockBinding(program, uniformBlockIndex, uniformBlockBinding)
}

func (R RenderingContext) Viewport(x, y, width, height int32) {
	defer R.checkError()
	R.RenderingContext.Viewport(x, y, width, height)
}

func (R RenderingContext) Scissor(x, y, width, height int32) {
	defer R.checkError()
	R.RenderingContext.Scissor(x, y, width, height)
}

func (R RenderingContext) Clear(mask glone.Enum) {
	defer R.checkError()
	R.RenderingContext.Clear(mask)
}

func (R RenderingContext) ClearBufferfv(buffer glone.Enum, drawbuffer int32, values []float32) {
	defer R.checkError()
	R.RenderingContext.ClearBufferfv(buffer, drawbuffer, values)
}

func (R RenderingContext) ClearBufferiv(buffer glone.Enum, drawbuffer int32, values []int32) {
	defer R.checkError()
	R.RenderingContext.ClearBufferiv(buffer, drawbuffer, values)
}

func (R RenderingContext) ClearBufferuiv(buffer glone.Enum, drawbuffer int32, values []uint32) {
	defer R.checkError()
	R.RenderingContext.ClearBufferuiv(buffer, drawbuffer, values)
}

func (R RenderingContext) ClearBufferfi(buffer glone.Enum, drawbuffer int32, depth float32, stencil int32) {
	defer R.checkError()
	R.RenderingContext.ClearBufferfi(buffer, drawbuffer, depth, stencil)
}

func (R RenderingContext) BlendColor(red, green, blue, alpha float32) {
	defer R.checkError()
	R.RenderingContext.BlendColor(red, green, blue, alpha)
}

func (R RenderingContext) BlendEquation(mode glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BlendEquation(mode)
}

func (R RenderingContext) BlendEquationSeparate(modeRGB, modeAlpha glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BlendEquationSeparate(modeRGB, modeAlpha)
}

func (R RenderingContext) BlendFunc(sfactor, dfactor glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BlendFunc(sfactor, dfactor)
}

func (R RenderingContext) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha glone.Enum) {
	defer R.checkError()
	R.RenderingContext.BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func (R RenderingContext) ClearColor(red, green, blue, alpha float32) {
	defer R.checkError()
	R.RenderingContext.ClearColor(red, green, blue, alpha)
}

func (R RenderingContext) ClearDepth(depth float32) {
	defer R.checkError()
	R.RenderingContext.ClearDepth(depth)
}

func (R RenderingContext) ClearStencil(s int32) {
	defer R.checkError()
	R.RenderingContext.ClearStencil(s)
}

func (R RenderingContext) ColorMask(red, green, blue, alpha bool) {
	defer R.checkError()
	R.RenderingContext.ColorMask(red, green, blue, alpha)
}

func (R RenderingContext) CullFace(mode glone.Enum) {
	defer R.checkError()
	R.RenderingContext.CullFace(mode)
}

func (R RenderingContext) FrontFace(mode glone.Enum) {
	defer R.checkError()
	R.RenderingContext.FrontFace(mode)
}

func (R RenderingContext) DepthFunc(fun glone.Enum) {
	defer R.checkError()
	R.RenderingContext.DepthFunc(fun)
}

func (R RenderingContext) DepthMask(flag bool) {
	defer R.checkError()
	R.RenderingContext.DepthMask(flag)
}

func (R RenderingContext) DepthRange(zNear, zFar float32) {
	defer R.checkError()
	R.RenderingContext.DepthRange(zNear, zFar)
}

func (R RenderingContext) Disable(cap glone.Enum) {
	defer R.checkError()
	R.RenderingContext.Disable(cap)
}

func (R RenderingContext) Enable(cap glone.Enum) {
	defer R.checkError()
	R.RenderingContext.Enable(cap)
}

func (R RenderingContext) Hint(target, mode glone.Enum) {
	defer R.checkError()
	R.RenderingContext.Hint(target, mode)
}

func (R RenderingContext) LineWidth(width float32) {
	defer R.checkError()
	R.RenderingContext.LineWidth(width)
}

func (R RenderingContext) PixelStorei(pname glone.Enum, param int32) {
	defer R.checkError()
	R.RenderingContext.PixelStorei(pname, param)
}

func (R RenderingContext) PolygonOffset(factor, units float32) {
	defer R.checkError()
	R.RenderingContext.PolygonOffset(factor, units)
}

func (R RenderingContext) StencilFunc(fun glone.Enum, ref int32, mask uint32) {
	defer R.checkError()
	R.RenderingContext.StencilFunc(fun, ref, mask)
}

func (R RenderingContext) StencilFuncSeparate(face, fun glone.Enum, ref int32, mask uint32) {
	defer R.checkError()
	R.RenderingContext.StencilFuncSeparate(face, fun, ref, mask)
}

func (R RenderingContext) StencilMask(mask uint32) {
	defer R.checkError()
	R.RenderingContext.StencilMask(mask)
}

func (R RenderingContext) StencilMaskSeparate(face glone.Enum, mask uint32) {
	defer R.checkError()
	R.RenderingContext.StencilMaskSeparate(face, mask)
}

func (R RenderingContext) StencilOp(fail, zfail, zpass glone.Enum) {
	defer R.checkError()
	R.RenderingContext.StencilOp(fail, zfail, zpass)
}

func (R RenderingContext) StencilOpSeparate(face, fail, zfail, zpass glone.Enum) {
	defer R.checkError()
	R.RenderingContext.StencilOpSeparate(face, fail, zfail, zpass)
}

func (R RenderingContext) SampleCoverage(value float32, invert bool) {
	defer R.checkError()
	R.RenderingContext.SampleCoverage(value, invert)
}

func (R RenderingContext) ReadBuffer(src glone.Enum) {
	defer R.checkError()
	R.RenderingContext.ReadBuffer(src)
}

func (R RenderingContext) DrawBuffers(buffers []glone.Enum) {
	defer R.checkError()
	R.RenderingContext.DrawBuffers(buffers)
}

func (R RenderingContext) GetParameter(pname glone.Enum) any {
	defer R.checkError()
	return R.RenderingContext.GetParameter(pname)
}

func (R RenderingContext) Finish() {
	defer R.checkError()
	R.RenderingContext.Finish()
}

func (R RenderingContext) Flush() {
	defer R.checkError()
	R.RenderingContext.Flush()
}

func (R RenderingContext) DrawArrays(mode glone.Enum, first, count int32) {
	defer R.checkError()
	R.RenderingContext.DrawArrays(mode, first, count)
}

func (R RenderingContext) DrawElements(mode glone.Enum, count int32, typ glone.Enum, offset int32) {
	defer R.checkError()
	R.RenderingContext.DrawElements(mode, count, typ, offset)
}

func (R RenderingContext) DrawArraysInstanced(mode glone.Enum, first, count, instanceCount int32) {
	defer R.checkError()
	R.RenderingContext.DrawArraysInstanced(mode, first, count, instanceCount)
}

func (R RenderingContext) DrawElementsInstanced(mode glone.Enum, count int32, typ glone.Enum, offset, instanceCount int32) {
	defer R.checkError()
	R.RenderingContext.DrawElementsInstanced(mode, count, typ, offset, instanceCount)
}

func (R RenderingContext) DrawRangeElements(mode glone.Enum, start, end uint32, count int32, typ glone.Enum, offset int32) {
	defer R.checkError()
	R.RenderingContext.DrawRangeElements(mode, start, end, count, typ, offset)
}

var _ glone.RenderingContext = RenderingContext{}
