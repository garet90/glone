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

func castToAny(value js.Value) any {
	switch value.Type() {
	case js.TypeUndefined, js.TypeNull:
		return nil
	case js.TypeBoolean:
		return value.Bool()
	case js.TypeNumber:
		return value.Float()
	case js.TypeString:
		return value.String()
	default:
		return nil
	}
}

func makeByteArray(v []byte) js.Value {

}

func makeFloat32Array(v []float32) js.Value {

}

func makeInt32Array(v []int32) js.Value {

}

func makeUint32Array(v []uint32) js.Value {

}

func makeUintArray(v []uint) js.Value {

}

func makeEnumArray(v []glone.Enum) js.Value {

}

func makeStringArray(v []string) js.Value {

}

func getUintArray(v js.Value) []uint {

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

func (R *RenderingContext) IsSync(sync glone.Sync) bool {
	s := syncOrNil(sync)
	return R.call("isSync", s).Bool()
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

func (R *RenderingContext) DeleteSync(sync glone.Sync) {
	s := syncOrNil(sync)
	R.call("deleteSync", s)
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
	if src.IsNull() {
		return ""
	}
	return src.String()
}

func (R *RenderingContext) GetShaderParameter(shader glone.Shader, pname glone.Enum) any {
	s := shaderOrNil(shader)
	return castToAny(R.call("getShaderParameter", s, pname))
}

func (R *RenderingContext) GetShaderPrecisionFormat(shadertype, precisiontype glone.Enum) glone.ShaderPrecisionFormat {
	return ShaderPrecisionFormat{
		R.call("getShaderPrecisionFormat", shadertype, precisiontype),
	}
}

func (R *RenderingContext) GetShaderInfoLog(shader glone.Shader) string {
	s := shaderOrNil(shader)
	log := R.call("getShaderInfoLog", s)
	if log.IsNull() {
		return ""
	}
	return log.String()
}

func (R *RenderingContext) AttachShader(program glone.Program, shader glone.Shader) {
	p := programOrNil(program)
	s := shaderOrNil(shader)
	R.call("attachShader", p, s)
}

func (R *RenderingContext) DetachShader(program glone.Program, shader glone.Shader) {
	p := programOrNil(program)
	s := shaderOrNil(shader)
	R.call("detachShader", p, s)
}

func (R *RenderingContext) LinkProgram(program glone.Program) {
	p := programOrNil(program)
	R.call("linkProgram", p)
}

func (R *RenderingContext) ValidateProgram(program glone.Program) {
	p := programOrNil(program)
	R.call("validateProgram", p)
}

func (R *RenderingContext) UseProgram(program glone.Program) {
	p := programOrNil(program)
	R.call("useProgram", p)
}

func (R *RenderingContext) GetActiveAttrib(program glone.Program, index uint) glone.ActiveInfo {
	p := programOrNil(program)
	return ActiveInfo{
		R.call("getActiveAttrib", p, index),
	}
}

func (R *RenderingContext) GetActiveUniform(program glone.Program, index uint) glone.ActiveInfo {
	p := programOrNil(program)
	return ActiveInfo{
		R.call("getActiveUniform", p, index),
	}
}

func (R *RenderingContext) GetAttachedShaders(program glone.Program) []glone.Shader {
	p := programOrNil(program)
	shaders := R.call("getAttachedShaders", p)
	shadersLength := shaders.Get("length").Int()
	out := make([]glone.Shader, shadersLength)
	for i := 0; i < shadersLength; i++ {
		out[i] = Shader{
			shaders.Index(i),
		}
	}
	return out
}

func (R *RenderingContext) GetAttribLocation(program glone.Program, name string) int {
	p := programOrNil(program)
	return R.call("getAttribLocation", p, name).Int()
}

func (R *RenderingContext) GetProgramInfoLog(program glone.Program) string {
	p := programOrNil(program)
	info := R.call("getProgramInfoLog", p)
	if info.IsNull() {
		return ""
	}
	return info.String()
}

func (R *RenderingContext) GetUniform(program glone.Program, location glone.UniformLocation) any {
	p := programOrNil(program)
	l := uniformLocationOrNil(location)
	return castToAny(R.call("getUniform", p, l))
}

func (R *RenderingContext) GetUniformLocation(program glone.Program, name string) glone.UniformLocation {
	p := programOrNil(program)
	return UniformLocation{
		R.call("getUniformLocation", p, name),
	}
}

func (R *RenderingContext) GetFragDataLocation(program glone.Program, name string) int {
	p := programOrNil(program)
	return R.call("getFragDataLocation", p, name).Int()
}

func (R *RenderingContext) GetProgramParameter(program glone.Program, pname glone.Enum) any {
	p := programOrNil(program)
	return castToAny(R.call("getProgramParameter", p, pname))
}

func (R *RenderingContext) Uniform1f(location glone.UniformLocation, x float64) {
	l := uniformLocationOrNil(location)
	R.call("uniform1f", l, x)
}

func (R *RenderingContext) Uniform2f(location glone.UniformLocation, x, y float64) {
	l := uniformLocationOrNil(location)
	R.call("uniform2f", l, x, y)
}

func (R *RenderingContext) Uniform3f(location glone.UniformLocation, x, y, z float64) {
	l := uniformLocationOrNil(location)
	R.call("uniform3f", l, x, y, z)
}

func (R *RenderingContext) Uniform4f(location glone.UniformLocation, x, y, z, w float64) {
	l := uniformLocationOrNil(location)
	R.call("uniform4f", l, x, y, z, w)
}

func (R *RenderingContext) Uniform1i(location glone.UniformLocation, x int) {
	l := uniformLocationOrNil(location)
	R.call("uniform1i", l, x)
}

func (R *RenderingContext) Uniform2i(location glone.UniformLocation, x, y int) {
	l := uniformLocationOrNil(location)
	R.call("uniform2i", l, x, y)
}

func (R *RenderingContext) Uniform3i(location glone.UniformLocation, x, y, z int) {
	l := uniformLocationOrNil(location)
	R.call("uniform3i", l, x, y, z)
}

func (R *RenderingContext) Uniform4i(location glone.UniformLocation, x, y, z, w int) {
	l := uniformLocationOrNil(location)
	R.call("uniform4i", l, x, y, z, w)
}

func (R *RenderingContext) Uniform1ui(location glone.UniformLocation, v0 uint) {
	l := uniformLocationOrNil(location)
	R.call("uniform1ui", l, v0)
}

func (R *RenderingContext) Uniform2ui(location glone.UniformLocation, v0, v1 uint) {
	l := uniformLocationOrNil(location)
	R.call("uniform2ui", l, v0, v1)
}

func (R *RenderingContext) Uniform3ui(location glone.UniformLocation, v0, v1, v2 uint) {
	l := uniformLocationOrNil(location)
	R.call("uniform3ui", l, v0, v1, v2)
}

func (R *RenderingContext) Uniform4ui(location glone.UniformLocation, v0, v1, v2, v3 uint) {
	l := uniformLocationOrNil(location)
	R.call("uniform4ui", l, v0, v1, v2, v3)
}

func (R *RenderingContext) Uniform1fv(location glone.UniformLocation, v []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniform1fv", l, makeFloat32Array(v))
}

func (R *RenderingContext) Uniform2fv(location glone.UniformLocation, v []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniform2fv", l, makeFloat32Array(v))
}

func (R *RenderingContext) Uniform3fv(location glone.UniformLocation, v []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniform3fv", l, makeFloat32Array(v))
}

func (R *RenderingContext) Uniform4fv(location glone.UniformLocation, v []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniform4fv", l, makeFloat32Array(v))
}

func (R *RenderingContext) Uniform1iv(location glone.UniformLocation, v []int32) {
	l := uniformLocationOrNil(location)
	R.call("uniform1iv", l, makeInt32Array(v))
}

func (R *RenderingContext) Uniform2iv(location glone.UniformLocation, v []int32) {
	l := uniformLocationOrNil(location)
	R.call("uniform2iv", l, makeInt32Array(v))
}

func (R *RenderingContext) Uniform3iv(location glone.UniformLocation, v []int32) {
	l := uniformLocationOrNil(location)
	R.call("uniform3iv", l, makeInt32Array(v))
}

func (R *RenderingContext) Uniform4iv(location glone.UniformLocation, v []int32) {
	l := uniformLocationOrNil(location)
	R.call("uniform4iv", l, makeInt32Array(v))
}

func (R *RenderingContext) Uniform1uiv(location glone.UniformLocation, data []uint32) {
	l := uniformLocationOrNil(location)
	R.call("uniform1uiv", l, makeUint32Array(data))
}

func (R *RenderingContext) Uniform2uiv(location glone.UniformLocation, data []uint32) {
	l := uniformLocationOrNil(location)
	R.call("uniform2uiv", l, makeUint32Array(data))
}

func (R *RenderingContext) Uniform3uiv(location glone.UniformLocation, data []uint32) {
	l := uniformLocationOrNil(location)
	R.call("uniform3uiv", l, makeUint32Array(data))
}

func (R *RenderingContext) Uniform4uiv(location glone.UniformLocation, data []uint32) {
	l := uniformLocationOrNil(location)
	R.call("uniform4uiv", l, makeUint32Array(data))
}

func (R *RenderingContext) UniformMatrix2fv(location glone.UniformLocation, transpose bool, value []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix2fv", l, transpose, makeFloat32Array(value))
}

func (R *RenderingContext) UniformMatrix3fv(location glone.UniformLocation, transpose bool, value []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix3fv", l, transpose, makeFloat32Array(value))
}

func (R *RenderingContext) UniformMatrix4fv(location glone.UniformLocation, transpose bool, value []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix4fv", l, transpose, makeFloat32Array(value))
}

func (R *RenderingContext) UniformMatrix3x2fv(location glone.UniformLocation, transpose bool, data []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix3x2fv", l, transpose, makeFloat32Array(data))
}

func (R *RenderingContext) UniformMatrix4x2fv(location glone.UniformLocation, transpose bool, data []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix4x2fv", l, transpose, makeFloat32Array(data))
}

func (R *RenderingContext) UniformMatrix2x3fv(location glone.UniformLocation, transpose bool, data []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix2x3fv", l, transpose, makeFloat32Array(data))
}

func (R *RenderingContext) UniformMatrix4x3fv(location glone.UniformLocation, transpose bool, data []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix4x3fv", l, transpose, makeFloat32Array(data))
}

func (R *RenderingContext) UniformMatrix2x4fv(location glone.UniformLocation, transpose bool, data []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix2x4fv", l, transpose, makeFloat32Array(data))
}

func (R *RenderingContext) UniformMatrix3x4fv(location glone.UniformLocation, transpose bool, data []float32) {
	l := uniformLocationOrNil(location)
	R.call("uniformMatrix2x4fv", l, transpose, makeFloat32Array(data))
}

func (R *RenderingContext) EnableVertexAttribArray(index uint) {
	R.call("enableVertexAttribArray", index)
}

func (R *RenderingContext) DisableVertexAttribArray(index uint) {
	R.call("disableVertexAttribArray", index)
}

func (R *RenderingContext) VertexAttrib1f(index uint, x float64) {
	R.call("vertexAttrib1f", index, x)
}

func (R *RenderingContext) VertexAttrib2f(index uint, x, y float64) {
	R.call("vertexAttrib2f", index, x, y)
}

func (R *RenderingContext) VertexAttrib3f(index uint, x, y, z float64) {
	R.call("vertexAttrib3f", index, x, y, z)
}

func (R *RenderingContext) VertexAttrib4f(index uint, x, y, z, w float64) {
	R.call("vertexAttrib4f", index, x, y, z, w)
}

func (R *RenderingContext) VertexAttrib1fv(index uint, values []float32) {
	R.call("vertexAttrib1fv", index, makeFloat32Array(values))
}

func (R *RenderingContext) VertexAttrib2fv(index uint, values []float32) {
	R.call("vertexAttrib2fv", index, makeFloat32Array(values))
}

func (R *RenderingContext) VertexAttrib3fv(index uint, values []float32) {
	R.call("vertexAttrib3fv", index, makeFloat32Array(values))
}

func (R *RenderingContext) VertexAttrib4fv(index uint, values []float32) {
	R.call("vertexAttrib4fv", index, makeFloat32Array(values))
}

func (R *RenderingContext) VertexAttribI4i(index uint, x, y, z, w int) {
	R.call("vertexAttribI4i", index, x, y, z, w)
}

func (R *RenderingContext) VertexAttribI4ui(index, x, y, z, w uint) {
	R.call("vertexAttribI4ui", index, x, y, z, w)
}

func (R *RenderingContext) VertexAttribI4iv(index uint, values []int32) {
	R.call("vertexAttribI4iv", index, makeInt32Array(values))
}

func (R *RenderingContext) VertexAttribI4uiv(index uint, values []uint32) {
	R.call("vertexAttribI4uiv", index, makeUint32Array(values))
}

func (R *RenderingContext) VertexAttribPointer(index uint, size int, typ glone.Enum, normalized bool, stride, offset int) {
	R.call("vertexAttribPointer", index, size, typ, normalized, stride, offset)
}

func (R *RenderingContext) VertexAttribIPointer(index uint, size int, typ glone.Enum, stride, offset int) {
	R.call("vertexAttribIPointer", index, size, typ, stride, offset)
}

func (R *RenderingContext) VertexAttribDivisor(index, divisor uint) {
	R.call("vertexAttribDivisor", index, divisor)
}

func (R *RenderingContext) GetVertexAttrib(index uint, pname glone.Enum) any {
	return castToAny(R.call("getVertexAttrib", index, pname))
}

func (R *RenderingContext) GetVertexAttribOffset(index uint, pname glone.Enum) int {
	return R.call("getVertexAttribOffset", index, pname).Int()
}

func (R *RenderingContext) FramebufferRenderbuffer(target, attachment, renderbuffertarget glone.Enum, renderbuffer glone.Renderbuffer) {
	r := renderbufferOrNil(renderbuffer)
	R.call("framebufferRenderbuffer", target, attachment, renderbuffertarget, r)
}

func (R *RenderingContext) FramebufferTexture2D(target, attachment, textarget glone.Enum, texture glone.Texture, level int) {
	t := textureOrNil(texture)
	R.call("framebufferTexture2D", target, attachment, textarget, t, level)
}

func (R *RenderingContext) CheckFramebufferStatus(target glone.Enum) glone.Enum {
	return glone.Enum(R.call("checkFramebufferStatus", target).Int())
}

func (R *RenderingContext) GetFramebufferAttachmentParameter(target, attachment, pname glone.Enum) any {
	return castToAny(R.call("getFramebufferAttachmentParameter", target, attachment, pname))
}

func (R *RenderingContext) BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask, filter glone.Enum) {
	R.call("blitFramebuffer", srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter)
}

func (R *RenderingContext) FramebufferTextureLayer(target, attachment glone.Enum, texture glone.Texture, level, layer int) {
	t := textureOrNil(texture)
	R.call("framebufferTextureLayer", target, attachment, t, level, layer)
}

func (R *RenderingContext) InvalidateFramebuffer(target glone.Enum, attachments []glone.Enum) {
	R.call("invalidateFramebuffer", target, makeEnumArray(attachments))
}

func (R *RenderingContext) InvalidateSubFramebuffer(target glone.Enum, attachments []glone.Enum, x, y, width, height int) {
	R.call("invalidateSubFramebuffer", target, makeEnumArray(attachments), x, y, width, height)
}

func (R *RenderingContext) ReadBuffer(src glone.Enum) {
	R.call("readBuffer", src)
}

func (R *RenderingContext) RenderbufferStorage(target, internalformat glone.Enum, width, height int) {
	R.call("renderbufferStorage", target, internalformat, width, height)
}

func (R *RenderingContext) RenderbufferStorageMultisample(target glone.Enum, samples int, internalformat glone.Enum, width, height int) {
	R.call("renderbufferStorageMultisample", target, samples, internalformat, width, height)
}

func (R *RenderingContext) GetRenderbufferParameter(target, pname glone.Enum) any {
	return castToAny(R.call("getRenderbufferParameter", target, pname))
}

func (R *RenderingContext) GetInternalFormatParameter(target, internalformat, pname glone.Enum) any {
	return castToAny(R.call("getInternalFormatParameter", target, internalformat, pname))
}

func (R *RenderingContext) ReadPixelsOff(x, y, width, height int, format, typ glone.Enum, offset int) {
	R.call("readPixels", x, y, width, height, format, typ, offset)
}

func (R *RenderingContext) ReadPixelsPix(x, y, width, height int, format, typ glone.Enum, dstData []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) DrawBuffers(buffers []glone.Enum) {
	R.call("drawBuffers", makeEnumArray(buffers))
}

func (R *RenderingContext) ClearBufferfv(buffer glone.Enum, drawbuffer int, values []float32) {
	R.call("clearBufferfv", buffer, drawbuffer, makeFloat32Array(values))
}

func (R *RenderingContext) ClearBufferiv(buffer glone.Enum, drawbuffer int, values []int32) {
	R.call("clearBufferiv", buffer, drawbuffer, makeInt32Array(values))
}

func (R *RenderingContext) ClearBufferuiv(buffer glone.Enum, drawbuffer int, values []uint32) {
	R.call("clearBufferuiv", buffer, drawbuffer, makeUint32Array(values))
}

func (R *RenderingContext) ClearBufferfi(buffer glone.Enum, drawbuffer int, depth float64, stencil int) {
	R.call("clearBufferfi", buffer, drawbuffer, depth, stencil)
}

func (R *RenderingContext) BufferDataSize(target glone.Enum, size int, usage glone.Enum) {
	R.call("bufferData", target, size, usage)
}

func (R *RenderingContext) BufferDataSrc(target glone.Enum, data glone.BufferSource, usage glone.Enum) {
	d := bufferSourceOrNil(data)
	R.call("bufferData", target, d, usage)
}

func (R *RenderingContext) BufferDataPix(target glone.Enum, srcData []byte, usage glone.Enum) {
	R.call("bufferData", target, makeByteArray(srcData), usage)
}

func (R *RenderingContext) BufferSubDataSrc(target glone.Enum, offset int, data glone.BufferSource) {
	d := bufferSourceOrNil(data)
	R.call("bufferSubData", target, offset, d)
}

func (R *RenderingContext) BufferSubDataPix(target glone.Enum, dstByteOffset int, srcData []byte) {
	R.call("bufferSubData", target, dstByteOffset, makeByteArray(srcData))
}

func (R *RenderingContext) GetBufferParameter(target, pname glone.Enum) any {
	return castToAny(R.call("getBufferParameter", target, pname))
}

func (R *RenderingContext) CopyBufferSubData(readTarget, writeTarget glone.Enum, readOffset, writeOffset, size int) {
	R.call("copyBufferSubData", readTarget, writeTarget, readOffset, writeOffset, size)
}

func (R *RenderingContext) GetBufferSubData(target glone.Enum, srcByteOffset int, dstBuffer []byte) {
	//TODO implement me
	panic("implement me")
}

func (R *RenderingContext) ActiveTexture(texture glone.Enum) {
	R.call("activeTexture", texture)
}

func (R *RenderingContext) TexParameterf(target, pname glone.Enum, param float64) {
	R.call("texParameterf", target, pname, param)
}

func (R *RenderingContext) TexParameteri(target, pname glone.Enum, param int) {
	R.call("texParameteri", target, pname, param)
}

func (R *RenderingContext) GetTexParameter(target, pname glone.Enum) any {
	return castToAny(R.call("getTexParameter", target, pname))
}

func (R *RenderingContext) GenerateMipmap(target glone.Enum) {
	R.call("generateMipmap", target)
}

func (R *RenderingContext) CopyTexImage2D(target glone.Enum, level int, internalformat glone.Enum, x, y, width, height, border int) {
	R.call("copyTexImage2D", target, level, internalformat, x, y, width, height, border)
}

func (R *RenderingContext) CopyTexSubImage2D(target glone.Enum, level int, xoffset, yoffset, x, y, width, height int) {
	R.call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, width, height)
}

func (R *RenderingContext) TexStorage2D(target glone.Enum, levels int, internalformat glone.Enum, width, height int) {
	R.call("texStorage2D", target, levels, internalformat, width, height)
}

func (R *RenderingContext) TexStorage3D(target glone.Enum, levels int, internalformat glone.Enum, width, height, depth int) {
	R.call("texStorage3D", target, levels, internalformat, width, height, depth)
}

func (R *RenderingContext) TexImage2DPbo(target glone.Enum, level, internalformat, width, height, border int, format, typ glone.Enum, pboOffset int) {
	R.call("texImage2D", target, level, internalformat, width, height, border, format, typ, pboOffset)
}

func (R *RenderingContext) TexImage2DSrc(target glone.Enum, level, internalformat, width, height, border int, format, typ glone.Enum, source glone.TexImageSource) {
	s := texImageSourceOrNil(source)
	R.call("texImage2D", target, level, internalformat, width, height, border, format, typ, s)
}

func (R *RenderingContext) TexImage2DPix(target glone.Enum, level, internalformat, width, height, border int, format, typ glone.Enum, srcData []byte) {
	R.call("texImage2D", target, level, internalformat, width, height, border, format, typ, makeByteArray(srcData))
}

func (R *RenderingContext) TexSubImage2DPbo(target glone.Enum, level, xoffset, yoffset, width, height int, format, typ glone.Enum, pboOffset int) {
	R.call("texSubImage2D", target, level, xoffset, yoffset, width, height, format, typ, pboOffset)
}

func (R *RenderingContext) TexSubImage2DSrc(target glone.Enum, level, xoffset, yoffset, width, height int, format, typ glone.Enum, source glone.TexImageSource) {
	s := texImageSourceOrNil(source)
	R.call("texSubImage2D", target, level, xoffset, yoffset, width, height, format, typ, s)
}

func (R *RenderingContext) TexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int, format, typ glone.Enum, srcData []byte) {
	R.call("texSubImage2D", target, level, xoffset, yoffset, width, height, format, typ, makeByteArray(srcData))
}

func (R *RenderingContext) TexImage3DPbo(target glone.Enum, level, internalformat, width, height, depth, border int, format, typ glone.Enum, pboOffset int) {
	R.call("texImage3D", target, level, internalformat, width, height, depth, border, format, typ, pboOffset)
}

func (R *RenderingContext) TexImage3DSrc(target glone.Enum, level, internalformat, width, height, depth, border int, format, typ glone.Enum, source glone.TexImageSource) {
	s := texImageSourceOrNil(source)
	R.call("texImage3D", target, level, internalformat, width, height, depth, border, format, typ, s)
}

func (R *RenderingContext) TexImage3DPix(target glone.Enum, level, internalformat, width, height, depth, border int, format, typ glone.Enum, srcData []byte) {
	R.call("texImage3D", target, level, internalformat, width, height, depth, border, format, typ, makeByteArray(srcData))
}

func (R *RenderingContext) TexSubImage3DPbo(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ glone.Enum, pboOffset int) {
	R.call("texSubImage3D", target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, pboOffset)
}

func (R *RenderingContext) TexSubImage3DSrc(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ glone.Enum, source glone.TexImageSource) {
	s := texImageSourceOrNil(source)
	R.call("texSubImage3D", target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, s)
}

func (R *RenderingContext) TexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ glone.Enum, srcData []byte) {
	R.call("texSubImage3D", target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, makeByteArray(srcData))
}

func (R *RenderingContext) CopyTexSubImage3D(target glone.Enum, level, xoffset, yoffset, zoffset, x, y, width, height int) {
	R.call("copyTexSubImage3D", target, level, xoffset, yoffset, zoffset, x, y, width, height)
}

func (R *RenderingContext) CompressedTexImage2DOff(target glone.Enum, level int, internalformat glone.Enum, width, height, border, imageSize, offset int) {
	R.call("compressedTexImage2D", target, level, internalformat, width, height, border, imageSize, offset)
}

func (R *RenderingContext) CompressedTexImage2DPix(target glone.Enum, level int, internalformat glone.Enum, width, height, border int, srcData []byte) {
	R.call("compressedTexImage2D", target, level, internalformat, width, height, border, makeByteArray(srcData))
}

func (R *RenderingContext) CompressedTexSubImage2DOff(target glone.Enum, level, xoffset, yoffset, width, height int, format glone.Enum, imageSize, offset int) {
	R.call("compressedTexSubImage2D", target, level, xoffset, yoffset, width, height, format, imageSize, offset)
}

func (R *RenderingContext) CompressedTexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int, format glone.Enum, srcData []byte) {
	R.call("compressedTexSubImage2D", target, level, xoffset, yoffset, width, height, format, makeByteArray(srcData))
}

func (R *RenderingContext) CompressedTexImage3D(target glone.Enum, level int, internalformat glone.Enum, width, height, depth, border, imageSize, offset int) {
	R.call("compressedTexImage3D", target, level, internalformat, width, height, depth, border, imageSize, offset)
}

func (R *RenderingContext) CompressedTexImage3DPix(target glone.Enum, level int, internalformat glone.Enum, width, height, depth, border int, srcData []byte) {
	R.call("compressedTexImage3D", target, level, internalformat, width, height, depth, border, makeByteArray(srcData))
}

func (R *RenderingContext) CompressedTexSubImage3D(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format glone.Enum, imageSize, offset int) {
	R.call("compressedTexSubImage3D", target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, offset)
}

func (R *RenderingContext) CompressedTexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format glone.Enum, srcData []byte) {
	R.call("compressedTexSubImage3D", target, level, xoffset, yoffset, zoffset, width, height, depth, format, makeByteArray(srcData))
}

func (R *RenderingContext) BeginQuery(target glone.Enum, query glone.Query) {
	q := queryOrNil(query)
	R.call("beginQuery", target, q)
}

func (R *RenderingContext) EndQuery(target glone.Enum) {
	R.call("endQuery", target)
}

func (R *RenderingContext) GetQuery(target, pname glone.Enum) glone.Query {
	return Query{
		R.call("getQuery", target, pname),
	}
}

func (R *RenderingContext) GetQueryParameter(query glone.Query, pname glone.Enum) any {
	q := queryOrNil(query)
	return castToAny(R.call("getQueryParameter", q, pname))
}

func (R *RenderingContext) SamplerParameteri(sampler glone.Sampler, pname glone.Enum, param int) {
	s := samplerOrNil(sampler)
	R.call("samplerParameteri", s, pname, param)
}

func (R *RenderingContext) SamplerParameterf(sampler glone.Sampler, pname glone.Enum, param float64) {
	s := samplerOrNil(sampler)
	R.call("samplerParameterf", s, pname, param)
}

func (R *RenderingContext) GetSamplerParameter(sampler glone.Sampler, pname glone.Enum) any {
	s := samplerOrNil(sampler)
	return castToAny(R.call("getSamplerParameter", s, pname))
}

func (R *RenderingContext) FenceSync(condition, flags glone.Enum) glone.Sync {
	return Sync{
		R.call("fenceSync", condition, flags),
	}
}

func (R *RenderingContext) ClientWaitSync(sync glone.Sync, flags glone.Enum, timeout uint) {
	s := syncOrNil(sync)
	R.call("clientWaitSync", s, flags, timeout)
}

func (R *RenderingContext) WaitSync(sync glone.Sync, flags glone.Enum, timeout int) {
	s := syncOrNil(sync)
	R.call("waitSync", s, flags, timeout)
}

func (R *RenderingContext) GetSyncParameter(sync glone.Sync, pname glone.Enum) any {
	s := syncOrNil(sync)
	return castToAny(R.call("getSyncParameter", s, pname))
}

func (R *RenderingContext) BeginTransformFeedback(primitiveMode glone.Enum) {
	R.call("beginTransformFeedback", primitiveMode)
}

func (R *RenderingContext) EndTransformFeedback() {
	R.call("endTransformFeedback")
}

func (R *RenderingContext) TransformFeedbackVaryings(program glone.Program, varyings []string, bufferMode glone.Enum) {
	p := programOrNil(program)
	R.call("transformFeedbackVaryings", p, makeStringArray(varyings), bufferMode)
}

func (R *RenderingContext) GetTransformFeedbackVarying(program glone.Program, index uint) glone.ActiveInfo {
	p := programOrNil(program)
	return ActiveInfo{
		R.call("getTransformFeedbackVarying", p, index),
	}
}

func (R *RenderingContext) PauseTransformFeedback() {
	R.call("pauseTransformFeedback")
}

func (R *RenderingContext) ResumeTransformFeedback() {
	R.call("resumeTransformFeedback")
}

func (R *RenderingContext) GetIndexedParameter(target glone.Enum, index uint) any {
	return castToAny(R.call("getIndexedParameter", target, index))
}

func (R *RenderingContext) GetUniformIndices(program glone.Program, uniformNames []string) []uint {
	p := programOrNil(program)
	return getUintArray(R.call("getUniformIndices", p, makeStringArray(uniformNames)))
}

func (R *RenderingContext) GetActiveUniforms(program glone.Program, uniformIndices []uint, pname glone.Enum) any {
	p := programOrNil(program)
	return castToAny(R.call("getActiveUniforms", p, makeUintArray(uniformIndices), pname))
}

func (R *RenderingContext) GetUniformBlockIndex(program glone.Program, uniformBlockName string) uint {
	p := programOrNil(program)
	return uint(R.call("getUniformBlockIndex", p, uniformBlockName).Int())
}

func (R *RenderingContext) GetActiveUniformBlockParameter(program glone.Program, uniformBlockIndex uint, pname glone.Enum) any {
	p := programOrNil(program)
	return castToAny(R.call("getActiveUniformBlockParameter", p, uniformBlockIndex, pname))
}

func (R *RenderingContext) GetActiveUniformBlockName(program glone.Program, uniformBlockIndex uint) string {
	p := programOrNil(program)
	return R.call("getActiveUniformBlockName", p, uniformBlockIndex).String()
}

func (R *RenderingContext) UniformBlockBinding(program glone.Program, uniformBlockIndex, uniformBlockBinding uint) {
	p := programOrNil(program)
	R.call("uniformBlockBinding", p, uniformBlockIndex, uniformBlockBinding)
}

func (R *RenderingContext) Viewport(x, y, width, height int) {
	R.call("viewport", x, y, width, height)
}

func (R *RenderingContext) Scissor(x, y, width, height int) {
	R.call("scissor", x, y, width, height)
}

func (R *RenderingContext) BlendColor(red, green, blue, alpha float64) {
	R.call("blendColor", red, green, blue, alpha)
}

func (R *RenderingContext) BlendEquation(mode glone.Enum) {
	R.call("blendEquation", mode)
}

func (R *RenderingContext) BlendEquationSeparate(modeRGB, modeAlpha glone.Enum) {
	R.call("blendEquationSeparate", modeRGB, modeAlpha)
}

func (R *RenderingContext) BlendFunc(sfactor, dfactor glone.Enum) {
	R.call("blendFunc", sfactor, dfactor)
}

func (R *RenderingContext) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha glone.Enum) {
	R.call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func (R *RenderingContext) Clear(mask glone.Enum) {
	R.call("clear", mask)
}

func (R *RenderingContext) ClearColor(red, green, blue, alpha float64) {
	R.call("clearColor", red, green, blue, alpha)
}

func (R *RenderingContext) ClearDepth(depth float64) {
	R.call("clearDepth", depth)
}

func (R *RenderingContext) ClearStencil(s int) {
	R.call("clearStencil", s)
}

func (R *RenderingContext) ColorMask(red, green, blue, alpha bool) {
	R.call("colorMask", red, green, blue, alpha)
}

func (R *RenderingContext) CullFace(mode glone.Enum) {
	R.call("cullFace", mode)
}

func (R *RenderingContext) FrontFace(mode glone.Enum) {
	R.call("frontFace", mode)
}

func (R *RenderingContext) DepthFunc(fun glone.Enum) {
	R.call("depthFunc", fun)
}

func (R *RenderingContext) DepthMask(flag bool) {
	R.call("depthMask", flag)
}

func (R *RenderingContext) DepthRange(zNear, zFar float64) {
	R.call("depthRange", zNear, zFar)
}

func (R *RenderingContext) Disable(cap glone.Enum) {
	R.call("disable", cap)
}

func (R *RenderingContext) Enable(cap glone.Enum) {
	R.call("enable", cap)
}

func (R *RenderingContext) Hint(target, mode glone.Enum) {
	R.call("hint", target, mode)
}

func (R *RenderingContext) LineWidth(width float64) {
	R.call("lineWidth", width)
}

func (R *RenderingContext) PixelStorei(pname glone.Enum, param int) {
	R.call("pixelStorei", pname, param)
}

func (R *RenderingContext) PolygonOffset(factor, units float64) {
	R.call("polygonOffset", factor, units)
}

func (R *RenderingContext) StencilFunc(fun glone.Enum, ref int, mask uint) {
	R.call("stencilFunc", fun, ref, mask)
}

func (R *RenderingContext) StencilFuncSeparate(face, fun glone.Enum, ref int, mask uint) {
	R.call("stencilFuncSeparate", face, fun, ref, mask)
}

func (R *RenderingContext) StencilMask(mask uint) {
	R.call("stencilMask", mask)
}

func (R *RenderingContext) StencilMaskSeparate(face glone.Enum, mask uint) {
	R.call("stencilMaskSeparate", face, mask)
}

func (R *RenderingContext) StencilOp(fail, zfail, zpass glone.Enum) {
	R.call("stencilOp", fail, zfail, zpass)
}

func (R *RenderingContext) StencilOpSeparate(face, fail, zfail, zpass glone.Enum) {
	R.call("stencilOpSeparate", face, fail, zfail, zpass)
}

func (R *RenderingContext) SampleCoverage(value float64, invert bool) {
	R.call("sampleCoverage", value, invert)
}

func (R *RenderingContext) GetParameter(pname glone.Enum) any {
	return castToAny(R.call("getParameter", pname))
}

func (R *RenderingContext) Finish() {
	R.call("finish")
}

func (R *RenderingContext) Flush() {
	R.call("flush")
}

func (R *RenderingContext) DrawArrays(mode glone.Enum, first, count int) {
	R.call("drawArrays", mode, first, count)
}

func (R *RenderingContext) DrawElements(mode glone.Enum, count int, typ glone.Enum, offset int) {
	R.call("drawElements", mode, count, typ, offset)
}

func (R *RenderingContext) DrawArraysInstanced(mode glone.Enum, first, count, instanceCount int) {
	R.call("drawArraysInstanced", mode, first, count, instanceCount)
}

func (R *RenderingContext) DrawElementsInstanced(mode glone.Enum, count int, typ glone.Enum, offset, instanceCount int) {
	R.call("drawElementsInstanced", mode, count, typ, offset, instanceCount)
}

func (R *RenderingContext) DrawRangeElements(mode glone.Enum, start, end uint, count int, typ glone.Enum, offset int) {
	R.call("drawRangeElements", mode, start, end, count, typ, offset)
}

func (R *RenderingContext) GetError() glone.Enum {
	return glone.Enum(R.call("getError").Int())
}

var _ glone.RenderingContext = (*RenderingContext)(nil)
