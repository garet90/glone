//go:build js && wasm

package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
)

type RenderingContext struct {
	value   js.Value
	methods map[string]js.Value
}

func NewRenderingContext(value js.Value) *RenderingContext {
	return &RenderingContext{
		value: value,
	}
}

func (R *RenderingContext) call(method string, args ...any) js.Value {
	if R.methods == nil {
		R.methods = make(map[string]js.Value)
	} else {
		m, ok := R.methods[method]
		if ok {
			return m.Invoke(args...)
		}
	}
	m := R.value.Get(method)
	R.methods[method] = m
	return m.Invoke(args...)
}

func (R *RenderingContext) CreateBuffer() glone.Buffer {
	v := R.call("createBuffer")
	if v.IsNull() {
		return nil
	}
	return Buffer{v}
}

func (R *RenderingContext) CreateFramebuffer() glone.Framebuffer {
	v := R.call("createFramebuffer")
	if v.IsNull() {
		return nil
	}
	return Framebuffer{v}
}

func (R *RenderingContext) CreateProgram() glone.Program {
	v := R.call("createProgram")
	if v.IsNull() {
		return nil
	}
	return Program{v}
}

func (R *RenderingContext) CreateRenderbuffer() glone.Renderbuffer {
	v := R.call("createRenderbuffer")
	if v.IsNull() {
		return nil
	}
	return Renderbuffer{v}
}

func (R *RenderingContext) CreateShader(typ glone.Enum) glone.Shader {
	v := R.call("createShader", typ)
	if v.IsNull() {
		return nil
	}
	return Renderbuffer{v}
}

func (R *RenderingContext) CreateTexture() glone.Texture {
	v := R.call("createTexture")
	if v.IsNull() {
		return nil
	}
	return Texture{v}
}

func (R *RenderingContext) CreateVertexArray() glone.VertexArray {
	v := R.call("createVertexArray")
	if v.IsNull() {
		return nil
	}
	return VertexArray{v}
}

func (R *RenderingContext) CreateTransformFeedback() glone.TransformFeedback {
	v := R.call("createTransformFeedback")
	if v.IsNull() {
		return nil
	}
	return TransformFeedback{v}
}

func (R *RenderingContext) CreateSampler() glone.Sampler {
	v := R.call("createSampler")
	if v.IsNull() {
		return nil
	}
	return Sampler{v}
}

func (R *RenderingContext) CreateQuery() glone.Query {
	v := R.call("createQuery")
	if v.IsNull() {
		return nil
	}
	return Query{v}
}

func (R *RenderingContext) BindAttribLocation(program glone.Program, index uint, name string) {
	p := programOrNil(program)
	R.call("bindAttribLocation", p, index, name)
}

func (R *RenderingContext) BindBuffer(target glone.Enum, buffer glone.Buffer) {
	b := bufferOrNil(buffer)
	R.call("bindBuffer", target, b)
}

func (R *RenderingContext) BindBufferBase(target glone.Enum, index uint, buffer glone.Buffer) {
	b := bufferOrNil(buffer)
	R.call("bindBufferBase", target, index, b)
}

func (R *RenderingContext) BindBufferRange(target glone.Enum, index uint, buffer glone.Buffer, offset, size int) {
	b := bufferOrNil(buffer)
	R.call("bindBufferRange", target, index, b, offset, size)
}

func (R *RenderingContext) BindFramebuffer(target glone.Enum, framebuffer glone.Framebuffer) {
	f := framebufferOrNil(framebuffer)
	R.call("bindFramebuffer", target, f)
}

func (R *RenderingContext) BindRenderbuffer(target glone.Enum, renderbuffer glone.Renderbuffer) {
	r := renderbufferOrNil(renderbuffer)
	R.call("bindRenderbuffer", target, r)
}

func (R *RenderingContext) BindTexture(target glone.Enum, texture glone.Texture) {
	t := textureOrNil(texture)
	R.call("bindTexture", target, t)
}

func (R *RenderingContext) BindVertexArray(array glone.VertexArray) {
	a := vertexArrayOrNil(array)
	R.call("bindVertexArray", a)
}

func (R *RenderingContext) BindTransformFeedback(target glone.Enum, tf glone.TransformFeedback) {
	t := transformFeedbackOrNil(tf)
	R.call("bindTransformFeedback", target, t)
}

func (R *RenderingContext) BindSampler(unit uint, sampler glone.Sampler) {
	s := samplerOrNil(sampler)
	R.call("bindSampler", unit, s)
}

func (R *RenderingContext) IsBuffer(buffer glone.Buffer) bool {
	b := bufferOrNil(buffer)
	return R.call("isBuffer", b).Bool()
}

func (R *RenderingContext) IsEnabled(cap glone.Enum) bool {
	return R.call("isEnabled", cap).Bool()
}

func (R *RenderingContext) IsFramebuffer(framebuffer glone.Framebuffer) bool {
	f := framebufferOrNil(framebuffer)
	return R.call("isFramebuffer", f).Bool()
}

func (R *RenderingContext) IsProgram(program glone.Program) bool {
	p := programOrNil(program)
	return R.call("isProgram", p).Bool()
}

func (R *RenderingContext) IsRenderbuffer(renderbuffer glone.Renderbuffer) bool {
	r := renderbufferOrNil(renderbuffer)
	return R.call("isRenderbuffer", r).Bool()
}

func (R *RenderingContext) IsShader(shader glone.Shader) bool {
	s := shaderOrNil(shader)
	return R.call("isShader", s).Bool()
}

func (R *RenderingContext) IsTexture(texture glone.Texture) bool {
	t := textureOrNil(texture)
	return R.call("isTexture", t).Bool()
}

func (R *RenderingContext) IsVertexArray(vertexArray glone.VertexArray) bool {
	v := vertexArrayOrNil(vertexArray)
	return R.call("isVertexArray", v).Bool()
}

func (R *RenderingContext) IsTransformFeedback(tf glone.TransformFeedback) bool {
	t := transformFeedbackOrNil(tf)
	return R.call("isTransformFeedback", t).Bool()
}

func (R *RenderingContext) IsSampler(sampler glone.Sampler) bool {
	s := samplerOrNil(sampler)
	return R.call("isSampler", s).Bool()
}

func (R *RenderingContext) IsQuery(query glone.Query) bool {
	q := queryOrNil(query)
	return R.call("isQuery", q).Bool()
}

func (R *RenderingContext) DeleteBuffer(buffer glone.Buffer) {
	b := bufferOrNil(buffer)
	R.call("deleteBuffer", b)
}

func (R *RenderingContext) DeleteFramebuffer(framebuffer glone.Framebuffer) {
	f := framebufferOrNil(framebuffer)
	R.call("deleteFramebuffer", f)
}

func (R *RenderingContext) DeleteProgram(program glone.Program) {
	p := programOrNil(program)
	R.call("deleteProgram", p)
}

func (R *RenderingContext) DeleteRenderbuffer(renderbuffer glone.Renderbuffer) {
	r := renderbufferOrNil(renderbuffer)
	R.call("deleteRenderbuffer", r)
}

func (R *RenderingContext) DeleteShader(shader glone.Shader) {
	s := shaderOrNil(shader)
	R.call("deleteShader", s)
}

func (R *RenderingContext) DeleteTexture(texture glone.Texture) {
	t := textureOrNil(texture)
	R.call("deleteTexture", t)
}

func (R *RenderingContext) DeleteVertexArray(vertexArray glone.VertexArray) {
	v := vertexArrayOrNil(vertexArray)
	R.call("deleteVertexArray", v)
}

func (R *RenderingContext) DeleteTransformFeedback(tf glone.TransformFeedback) {
	t := transformFeedbackOrNil(tf)
	R.call("deleteTransformFeedback", t)
}

func (R *RenderingContext) DeleteSampler(sampler glone.Sampler) {
	s := samplerOrNil(sampler)
	R.call("deleteSampler", s)
}

func (R *RenderingContext) DeleteQuery(query glone.Query) {
	q := queryOrNil(query)
	R.call("deleteQuery", q)
}

func (R *RenderingContext) ShaderSource(shader glone.Shader, source string) {
	s := shaderOrNil(shader)
	R.call("shaderSource", s, source)
}

func (R *RenderingContext) CompileShader(shader glone.Shader) {
	s := shaderOrNil(shader)
	R.call("compileShader", s)
}

func (R *RenderingContext) GetShaderSource(shader glone.Shader) string {
	s := shaderOrNil(shader)
	src := R.call("getShaderSource", s)
	if src.Truthy()
}

func (R *RenderingContext) GetShaderParameter(shader glone.Shader, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetShaderPrecisionFormat(shadertype, precisiontype glone.Enum) glone.ShaderPrecisionFormat {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetShaderInfoLog(shader glone.Shader) string {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) AttachShader(program glone.Program, shader glone.Shader) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ValidateProgram(program glone.Program) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UseProgram(program glone.Program) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetActiveAttrib(program glone.Program, index uint) glone.ActiveInfo {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetActiveUniform(program glone.Program, index uint) glone.ActiveInfo {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetAttachedShaders(program glone.Program) []glone.Shader {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetAttribLocation(program glone.Program, name string) int {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetProgramInfoLog(program glone.Program) string {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetUniform(program glone.Program, location glone.UniformLocation) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetUniformLocation(program glone.Program, name string) glone.UniformLocation {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetFragDataLocation(program glone.Program, name string) int {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetProgramParameter(program glone.Program, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform1f(location glone.UniformLocation, x float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform2f(location glone.UniformLocation, x, y float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform3f(location glone.UniformLocation, x, y, z float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform4f(location glone.UniformLocation, x, y, z, w float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform1i(location glone.UniformLocation, x int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform2i(location glone.UniformLocation, x, y int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform3i(location glone.UniformLocation, x, y, z int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform4i(location glone.UniformLocation, x, y, z, w int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform1ui(location glone.UniformLocation, v0 uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform2ui(location glone.UniformLocation, v0, v1 uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform3ui(location glone.UniformLocation, v0, v1, v2 uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform4ui(location glone.UniformLocation, v0, v1, v2, v3 uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform1fv(location glone.UniformLocation, v []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform2fv(location glone.UniformLocation, v []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform3fv(location glone.UniformLocation, v []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform4fv(location glone.UniformLocation, v []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform1iv(location glone.UniformLocation, v []int32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform2iv(location glone.UniformLocation, v []int32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform3iv(location glone.UniformLocation, v []int32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform4iv(location glone.UniformLocation, v []int32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform1uiv(location glone.UniformLocation, data []uint32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform2uiv(location glone.UniformLocation, data []uint32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform3uiv(location glone.UniformLocation, data []uint32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Uniform4uiv(location glone.UniformLocation, data []uint32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix2fv(location glone.UniformLocation, transpose bool, value []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix3fv(location glone.UniformLocation, transpose bool, value []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix4fv(location glone.UniformLocation, transpose bool, value []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix3x2fv(location glone.UniformLocation, transpose bool, data []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix4x2fv(location glone.UniformLocation, transpose bool, data []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix2x3fv(location glone.UniformLocation, transpose bool, data []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix4x3fv(location glone.UniformLocation, transpose bool, data []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix2x4fv(location glone.UniformLocation, transpose bool, data []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformMatrix3x4fv(location glone.UniformLocation, transpose bool, data []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) EnableVertexAttribArray(index uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DisableVertexAttribArray(index uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttrib1f(index uint, x float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttrib2f(index uint, x, y float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttrib3f(index uint, x, y, z float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttrib4f(index uint, x, y, z, w float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttrib1fv(index uint, values []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttrib2fv(index uint, values []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttrib3fv(index uint, values []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttrib4fv(index uint, values []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttribI4i(index uint, x, y, z, w int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttribI4iv(index uint, values []int32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttribI4ui(index, x, y, z, w uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttribI4uiv(index uint, values []uint32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttribPointer(index uint, size int, typ glone.Enum, normalized bool, stride, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttribIPointer(index uint, size int, typ glone.Enum, stride, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) VertexAttribDivisor(index, divisor uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetVertexAttrib(index uint, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetVertexAttribOffset(index uint, pname glone.Enum) int {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) FramebufferRenderbuffer(target, attachment, renderbuffertarget glone.Enum, renderbuffer glone.Renderbuffer) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) FramebufferTexture2D(target, attachment, textarget glone.Enum, texture glone.Texture, level int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CheckFramebufferStatus(target glone.Enum) glone.Enum {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetFramebufferAttachmentParameter(target, attachnemt, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask, filter glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) FramebufferTextureLayer(target, attachment glone.Enum, texture glone.Texture, level, layer int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) InvalidateFramebuffer(target glone.Enum, attachments []glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) InvalidateSubFramebuffer(target glone.Enum, attachments []glone.Enum, x, y, width, height int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ReadBuffer(src glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) RenderbufferStorage(target, internalformat glone.Enum, width, height int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) RenderbufferStorageMultisample(target glone.Enum, samples int, internalformat glone.Enum, width, height int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetRenderbufferParameter(target, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetInternalFormatParameter(target, internalformat, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ReadPixelsOff(x, y, width, height int, format, typ glone.Enum, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ReadPixelsPix(x, y, width, height int, format, typ glone.Enum, dstData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DrawBuffers(buffers []glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ClearBufferfv(buffer glone.Enum, drawbuffer int, values []float32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ClearBufferiv(buffer glone.Enum, drawbuffer int, values []int32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ClearBufferuiv(buffer glone.Enum, drawbuffer int, values []uint32) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ClearBufferfi(buffer glone.Enum, drawbuffer int, depth float64, stencil int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BufferDataSize(target glone.Enum, size int, usage glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BufferDataSrc(target glone.Enum, data glone.BufferSource, usage glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BufferDataPix(target glone.Enum, srcData []byte, usage glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BufferSubDataSrc(target glone.Enum, offset int, date glone.BufferSource) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BufferSubDataPix(target glone.Enum, dstByteOffset int, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetBufferParameter(target, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CopyBufferSubData(readTarget, writeTarget glone.Enum, readOffset, writeOffset, size int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetBufferSubData(target glone.Enum, srcByteOffset int, dstBuffer []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ActiveTexture(texture glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexParameterf(target, pname glone.Enum, param float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexParameteri(target, pname glone.Enum, param int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetTexParameter(target, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GenerateMipmap(target glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CopyTexImage2D(target glone.Enum, level int, internalFormat glone.Enum, x, y, width, height, border int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CopyTexSubImage2D(target glone.Enum, level int, xoffset, yoffset, x, y, width, height int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexStorage2D(target glone.Enum, levels int, internalformat glone.Enum, width, height int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexStorage3D(target glone.Enum, levels int, internalformat glone.Enum, width, height, depth int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexImage2DPbo(target glone.Enum, level, internalformat, width, height, border int, format, typ glone.Enum, pboOffset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexImage2DSrc(target glone.Enum, level, internalformat, width, height, border int, format, typ glone.Enum, source glone.TexImageSource) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexImage2DPix(target glone.Enum, level, internalformat, width, height, border int, format, typ glone.Enum, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexSubImage2DPbo(target glone.Enum, level, xoffset, yoffset, width, height int, format, typ glone.Enum, pboOffset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexSubImage2DSrc(target glone.Enum, level, xoffset, yoffset, width, height int, format, typ glone.Enum, source glone.TexImageSource) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int, format, typ glone.Enum, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexImage3DPbo(target glone.Enum, level, internalformat, width, height, depth, border int, format, typ glone.Enum, pboOffset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexImage3DSrc(target glone.Enum, level, internalformat, width, height, depth, border int, format, typ glone.Enum, source glone.TexImageSource) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexImage3DPix(target glone.Enum, level, internalformat, width, height, depth, border int, format, typ glone.Enum, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexSubImage3DPbo(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ glone.Enum, pboOffset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexSubImage3DSrc(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ glone.Enum, source glone.TexImageSource) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ glone.Enum, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CopyTexSubImage3D(target glone.Enum, level, xoffset, yoffset, zoffset, x, y, width, height int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CompressedTexImage2DOff(target glone.Enum, level int, internalformat glone.Enum, width, height, border, imageSize, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CompressedTexImage2DPix(target glone.Enum, level int, internalformat glone.Enum, width, height, border int, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CompressedTexSubImage2DOff(target glone.Enum, level, xoffset, yoffset, width, height int, format glone.Enum, imageSize, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CompressedTexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int, format glone.Enum, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CompressedTexImage3D(target glone.Enum, level int, internalformat glone.Enum, width, height, depth, border, imageSize, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CompressedTexImage3DSrc(target glone.Enum, level int, internalformat glone.Enum, width, height, depth, border int, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CompressedTexSubImage3D(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format glone.Enum, imageSize, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CompressedTexSubImage3DSrc(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format glone.Enum, srcData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BeginQuery(target glone.Enum, query glone.Query) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) EndQuery(target glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetQuery(target, pname glone.Enum) glone.Query {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetQueryParameter(query glone.Query, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) SamplerParameteri(sampler glone.Sampler, pname glone.Enum, param int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) SamplerParameterf(sampler glone.Sampler, pname glone.Enum, param float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetSamplerParameter(sampler glone.Sampler, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) FenceSync(condition, flags glone.Enum) glone.Sync {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) IsSync(sync glone.Sync) bool {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DeleteSync(sync glone.Sync) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ClientWaitSync(sync glone.Sync, flags glone.Enum, timeout uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) WaitSync(sync glone.Sync, flags glone.Enum, timeout int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetSyncParameter(sync glone.Sync, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BeginTransformFeedback(primitiveMode glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) EndTransformFeedback() {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) TransformFeedbackVaryings(program glone.Program, varyings []string, bufferMode glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetTransformFeedbackVarying(program glone.Program, index uint) glone.ActiveInfo {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) PauseTransformFeedback() {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ResumeTransformFeedback() {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetIndexedParameter(target glone.Enum, index uint) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetUniformIndices(program glone.Program, uniformNames []string) []uint {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetActiveUniforms(program glone.Program, uniformIndices []uint, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetUniformBlockIndex(program glone.Program, uniformBlockName string) uint {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetActiveUniformBlockParameter(program glone.Program, uniformBlockIndex uint, pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetActiveUniformBlockName(program glone.Program, uniformBlockIndex uint) string {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) UniformBlockBinding(program glone.Program, uniformBlockIndex, uniformBlockBinding uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Viewport(x, y, width, height int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Scissor(x, y, width, height int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BlendColor(red, green, blue, alpha float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BlendEquation(mode glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BlendEquationSeparate(modeRGB, modeAlpha glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BlendFunc(sfactor, dfactor glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Clear(mask glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ClearColor(red, green, blue, alpha float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ClearDepth(depth float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ClearStencil(s int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ColorMask(red, green, blue, alpha bool) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) CullFace(mode glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) FrontFace(mode glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DepthFunc(fun glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DepthMask(flag bool) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DepthRange(zNear, zFar float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DetachShader(program glone.Program, shader glone.Shader) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Disable(cap glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Enable(cap glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Hint(target, mode glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) LineWidth(width float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) LinkProgram(program glone.Program) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) PixelStorei(pname glone.Enum, param int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) PolygonOffset(factor, units float64) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) StencilFunc(fun glone.Enum, ref int, mask uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) StencilFuncSeparate(face, fun glone.Enum, ref int, mask uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) StencilMask(mask uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) StencilMaskSeparate(face glone.Enum, mask uint) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) StencilOp(fail, zfail, zpass glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) StencilOpSeparate(face, fail, zfail, zpass glone.Enum) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) SampleCoverage(value float64, invert bool) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetParameter(pname glone.Enum) any {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Finish() {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) Flush() {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DrawArrays(mode glone.Enum, first, count int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DrawElements(mode glone.Enum, count int, typ glone.Enum, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DrawArraysInstanced(mode glone.Enum, first, count, instanceCount int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DrawElementsInstanced(mode glone.Enum, count int, typ glone.Enum, offset, instanceCount int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DrawRangeElements(mode glone.Enum, start, end uint, count int, typ glone.Enum, offset int) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) GetError() glone.Enum {
	//TODO implement me
	panic("implement me")
}

var _ glone.RenderingContext = (*RenderingContext)(nil)
