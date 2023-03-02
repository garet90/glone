//go:build js && wasm

package webgl

import (
	"gfx.cafe/ghalliday1/glone"
	"syscall/js"
	"unsafe"
)

type RenderingContext struct {
	value   js.Value
	methods map[string]js.Value
}

func MakeRenderingContext(value js.Value) RenderingContext {
	return RenderingContext{
		value: value,
	}
}

func NewRenderingContext(value js.Value) *RenderingContext {
	rc := MakeRenderingContext(value)
	return &rc
}

var (
	valueGlobal = js.Global()

	arrayConstructor        = valueGlobal.Get("Array")
	uint8ArrayConstructor   = valueGlobal.Get("Uint8Array")
	float32ArrayConstructor = valueGlobal.Get("Float32Array")
	int32ArrayConstructor   = valueGlobal.Get("Int32Array")
	uint32ArrayConstructor  = valueGlobal.Get("Uint32Array")
)

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

func makeByteArrayLength(length int) js.Value {
	return uint8ArrayConstructor.New(length)
}

func makeByteArray(v []byte) js.Value {
	array := makeByteArrayLength(len(v))
	js.CopyBytesToJS(array, v)
	return array
}

func unsafeReslice[T, U any](v []T) []U {
	tSize := unsafe.Sizeof(*new(T))
	uSize := unsafe.Sizeof(*new(U))
	data := unsafe.SliceData(v)
	return unsafe.Slice((*U)(unsafe.Pointer(data)), (int(tSize)*len(v))/int(uSize))
}

func makeFloat32Array(v []float32) js.Value {
	array := makeByteArray(unsafeReslice[float32, byte](v))
	return float32ArrayConstructor.New(array.Get("buffer"))
}

func makeInt32Array(v []int32) js.Value {
	array := makeByteArray(unsafeReslice[int32, byte](v))
	return int32ArrayConstructor.New(array.Get("buffer"))
}

func makeUint32Array(v []uint32) js.Value {
	array := makeByteArray(unsafeReslice[uint32, byte](v))
	return uint32ArrayConstructor.New(array.Get("buffer"))
}

func makeUintArray(v []uint32) js.Value {
	array := arrayConstructor.New(len(v))
	for i, val := range v {
		array.SetIndex(i, val)
	}
	return array
}

func makeEnumArray(v []glone.Enum) js.Value {
	array := arrayConstructor.New(len(v))
	for i, val := range v {
		array.SetIndex(i, int32(val))
	}
	return array
}

func makeStringArray(v []string) js.Value {
	array := arrayConstructor.New(len(v))
	for i, val := range v {
		array.SetIndex(i, val)
	}
	return array
}

func getUintArray(v js.Value) []uint32 {
	length := v.Get("length").Int()
	array := make([]uint32, length)
	for i := 0; i < length; i++ {
		array[i] = uint32(v.Index(i).Int())
	}
	return array
}

func getShaderArray(v js.Value) []glone.Shader {
	length := v.Get("length").Int()
	array := make([]glone.Shader, length)
	for i := 0; i < length; i++ {
		array[i] = &Shader{
			v.Index(i),
		}
	}
	return array
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
	m = m.Call("bind", R.value)
	R.methods[method] = m
	return m.Invoke(args...)
}

func (R *RenderingContext) CreateBuffer() glone.Buffer {
	v := R.call("createBuffer")
	if v.IsNull() {
		return nil
	}
	return &Buffer{v}
}

func (R *RenderingContext) CreateFramebuffer() glone.Framebuffer {
	v := R.call("createFramebuffer")
	if v.IsNull() {
		return nil
	}
	return &Framebuffer{v}
}

func (R *RenderingContext) CreateProgram() glone.Program {
	v := R.call("createProgram")
	if v.IsNull() {
		return nil
	}
	return &Program{v}
}

func (R *RenderingContext) CreateRenderbuffer() glone.Renderbuffer {
	v := R.call("createRenderbuffer")
	if v.IsNull() {
		return nil
	}
	return &Renderbuffer{v}
}

func (R *RenderingContext) CreateShader(typ glone.Enum) glone.Shader {
	v := R.call("createShader", int32(typ))
	if v.IsNull() {
		return nil
	}
	return &Shader{v}
}

func (R *RenderingContext) CreateTexture() glone.Texture {
	v := R.call("createTexture")
	if v.IsNull() {
		return nil
	}
	return &Texture{v}
}

func (R *RenderingContext) CreateVertexArray() glone.VertexArray {
	v := R.call("createVertexArray")
	if v.IsNull() {
		return nil
	}
	return &VertexArray{v}
}

func (R *RenderingContext) CreateTransformFeedback() glone.TransformFeedback {
	v := R.call("createTransformFeedback")
	if v.IsNull() {
		return nil
	}
	return &TransformFeedback{v}
}

func (R *RenderingContext) CreateSampler() glone.Sampler {
	v := R.call("createSampler")
	if v.IsNull() {
		return nil
	}
	return &Sampler{v}
}

func (R *RenderingContext) CreateQuery() glone.Query {
	v := R.call("createQuery")
	if v.IsNull() {
		return nil
	}
	return &Query{v}
}

func (R *RenderingContext) BindAttribLocation(program glone.Program, index uint32, name string) {
	p := programOrNil(program)
	R.call("bindAttribLocation", p, index, name)
}

func (R *RenderingContext) BindBuffer(target glone.Enum, buffer glone.Buffer) {
	b := bufferOrNil(buffer)
	R.call("bindBuffer", int32(target), b)
}

func (R *RenderingContext) BindBufferBase(target glone.Enum, index uint32, buffer glone.Buffer) {
	b := bufferOrNil(buffer)
	R.call("bindBufferBase", int32(target), index, b)
}

func (R *RenderingContext) BindBufferRange(target glone.Enum, index uint32, buffer glone.Buffer, offset, size int32) {
	b := bufferOrNil(buffer)
	R.call("bindBufferRange", int32(target), index, b, offset, size)
}

func (R *RenderingContext) BindFramebuffer(target glone.Enum, framebuffer glone.Framebuffer) {
	f := framebufferOrNil(framebuffer)
	R.call("bindFramebuffer", int32(target), f)
}

func (R *RenderingContext) BindRenderbuffer(target glone.Enum, renderbuffer glone.Renderbuffer) {
	r := renderbufferOrNil(renderbuffer)
	R.call("bindRenderbuffer", int32(target), r)
}

func (R *RenderingContext) BindTexture(target glone.Enum, texture glone.Texture) {
	t := textureOrNil(texture)
	R.call("bindTexture", int32(target), t)
}

func (R *RenderingContext) BindVertexArray(array glone.VertexArray) {
	a := vertexArrayOrNil(array)
	R.call("bindVertexArray", a)
}

func (R *RenderingContext) BindTransformFeedback(target glone.Enum, tf glone.TransformFeedback) {
	t := transformFeedbackOrNil(tf)
	R.call("bindTransformFeedback", int32(target), t)
}

func (R *RenderingContext) BindSampler(unit uint32, sampler glone.Sampler) {
	s := samplerOrNil(sampler)
	R.call("bindSampler", unit, s)
}

func (R *RenderingContext) IsBuffer(buffer glone.Buffer) bool {
	b := bufferOrNil(buffer)
	return R.call("isBuffer", b).Bool()
}

func (R *RenderingContext) IsEnabled(cap glone.Enum) bool {
	return R.call("isEnabled", int32(cap)).Bool()
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
	return castToAny(R.call("getShaderParameter", s, int32(pname)))
}

func (R *RenderingContext) GetShaderPrecisionFormat(shadertype, precisiontype glone.Enum) glone.ShaderPrecisionFormat {
	return &ShaderPrecisionFormat{
		R.call("getShaderPrecisionFormat", int32(shadertype), int32(precisiontype)),
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

func (R *RenderingContext) GetActiveAttrib(program glone.Program, index uint32) glone.ActiveInfo {
	p := programOrNil(program)
	return &ActiveInfo{
		R.call("getActiveAttrib", p, index),
	}
}

func (R *RenderingContext) GetActiveUniform(program glone.Program, index uint32) glone.ActiveInfo {
	p := programOrNil(program)
	return &ActiveInfo{
		R.call("getActiveUniform", p, index),
	}
}

func (R *RenderingContext) GetAttachedShaders(program glone.Program) []glone.Shader {
	p := programOrNil(program)
	return getShaderArray(R.call("getAttachedShaders", p))
}

func (R *RenderingContext) GetAttribLocation(program glone.Program, name string) int32 {
	p := programOrNil(program)
	return int32(R.call("getAttribLocation", p, name).Int())
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
	return &UniformLocation{
		R.call("getUniformLocation", p, name),
	}
}

func (R *RenderingContext) GetFragDataLocation(program glone.Program, name string) int32 {
	p := programOrNil(program)
	return int32(R.call("getFragDataLocation", p, name).Int())
}

func (R *RenderingContext) GetProgramParameter(program glone.Program, pname glone.Enum) any {
	p := programOrNil(program)
	return castToAny(R.call("getProgramParameter", p, int32(pname)))
}

func (R *RenderingContext) Uniform1f(location glone.UniformLocation, x float32) {
	l := uniformLocationOrNil(location)
	R.call("uniform1f", l, x)
}

func (R *RenderingContext) Uniform2f(location glone.UniformLocation, x, y float32) {
	l := uniformLocationOrNil(location)
	R.call("uniform2f", l, x, y)
}

func (R *RenderingContext) Uniform3f(location glone.UniformLocation, x, y, z float32) {
	l := uniformLocationOrNil(location)
	R.call("uniform3f", l, x, y, z)
}

func (R *RenderingContext) Uniform4f(location glone.UniformLocation, x, y, z, w float32) {
	l := uniformLocationOrNil(location)
	R.call("uniform4f", l, x, y, z, w)
}

func (R *RenderingContext) Uniform1i(location glone.UniformLocation, x int32) {
	l := uniformLocationOrNil(location)
	R.call("uniform1i", l, x)
}

func (R *RenderingContext) Uniform2i(location glone.UniformLocation, x, y int32) {
	l := uniformLocationOrNil(location)
	R.call("uniform2i", l, x, y)
}

func (R *RenderingContext) Uniform3i(location glone.UniformLocation, x, y, z int32) {
	l := uniformLocationOrNil(location)
	R.call("uniform3i", l, x, y, z)
}

func (R *RenderingContext) Uniform4i(location glone.UniformLocation, x, y, z, w int32) {
	l := uniformLocationOrNil(location)
	R.call("uniform4i", l, x, y, z, w)
}

func (R *RenderingContext) Uniform1ui(location glone.UniformLocation, v0 uint32) {
	l := uniformLocationOrNil(location)
	R.call("uniform1ui", l, v0)
}

func (R *RenderingContext) Uniform2ui(location glone.UniformLocation, v0, v1 uint32) {
	l := uniformLocationOrNil(location)
	R.call("uniform2ui", l, v0, v1)
}

func (R *RenderingContext) Uniform3ui(location glone.UniformLocation, v0, v1, v2 uint32) {
	l := uniformLocationOrNil(location)
	R.call("uniform3ui", l, v0, v1, v2)
}

func (R *RenderingContext) Uniform4ui(location glone.UniformLocation, v0, v1, v2, v3 uint32) {
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

func (R *RenderingContext) EnableVertexAttribArray(index uint32) {
	R.call("enableVertexAttribArray", index)
}

func (R *RenderingContext) DisableVertexAttribArray(index uint32) {
	R.call("disableVertexAttribArray", index)
}

func (R *RenderingContext) VertexAttrib1f(index uint32, x float32) {
	R.call("vertexAttrib1f", index, x)
}

func (R *RenderingContext) VertexAttrib2f(index uint32, x, y float32) {
	R.call("vertexAttrib2f", index, x, y)
}

func (R *RenderingContext) VertexAttrib3f(index uint32, x, y, z float32) {
	R.call("vertexAttrib3f", index, x, y, z)
}

func (R *RenderingContext) VertexAttrib4f(index uint32, x, y, z, w float32) {
	R.call("vertexAttrib4f", index, x, y, z, w)
}

func (R *RenderingContext) VertexAttrib1fv(index uint32, values []float32) {
	R.call("vertexAttrib1fv", index, makeFloat32Array(values))
}

func (R *RenderingContext) VertexAttrib2fv(index uint32, values []float32) {
	R.call("vertexAttrib2fv", index, makeFloat32Array(values))
}

func (R *RenderingContext) VertexAttrib3fv(index uint32, values []float32) {
	R.call("vertexAttrib3fv", index, makeFloat32Array(values))
}

func (R *RenderingContext) VertexAttrib4fv(index uint32, values []float32) {
	R.call("vertexAttrib4fv", index, makeFloat32Array(values))
}

func (R *RenderingContext) VertexAttribI4i(index uint32, x, y, z, w int32) {
	R.call("vertexAttribI4i", index, x, y, z, w)
}

func (R *RenderingContext) VertexAttribI4ui(index, x, y, z, w uint32) {
	R.call("vertexAttribI4ui", index, x, y, z, w)
}

func (R *RenderingContext) VertexAttribI4iv(index uint32, values []int32) {
	R.call("vertexAttribI4iv", index, makeInt32Array(values))
}

func (R *RenderingContext) VertexAttribI4uiv(index uint32, values []uint32) {
	R.call("vertexAttribI4uiv", index, makeUint32Array(values))
}

func (R *RenderingContext) VertexAttribPointer(index uint32, size int32, typ glone.Enum, normalized bool, stride, offset int32) {
	R.call("vertexAttribPointer", index, size, int32(typ), normalized, stride, offset)
}

func (R *RenderingContext) VertexAttribIPointer(index uint32, size int32, typ glone.Enum, stride, offset int32) {
	R.call("vertexAttribIPointer", index, size, int32(typ), stride, offset)
}

func (R *RenderingContext) VertexAttribDivisor(index, divisor uint32) {
	R.call("vertexAttribDivisor", index, divisor)
}

func (R *RenderingContext) GetVertexAttrib(index uint32, pname glone.Enum) any {
	return castToAny(R.call("getVertexAttrib", index, int32(pname)))
}

func (R *RenderingContext) GetVertexAttribOffset(index uint32, pname glone.Enum) int32 {
	return int32(R.call("getVertexAttribOffset", index, int32(pname)).Int())
}

func (R *RenderingContext) FramebufferRenderbuffer(target, attachment, renderbuffertarget glone.Enum, renderbuffer glone.Renderbuffer) {
	r := renderbufferOrNil(renderbuffer)
	R.call("framebufferRenderbuffer", int32(target), int32(attachment), int32(renderbuffertarget), r)
}

func (R *RenderingContext) FramebufferTexture2D(target, attachment, textarget glone.Enum, texture glone.Texture, level int32) {
	t := textureOrNil(texture)
	R.call("framebufferTexture2D", int32(target), int32(attachment), int32(textarget), t, level)
}

func (R *RenderingContext) CheckFramebufferStatus(target glone.Enum) glone.Enum {
	return glone.Enum(R.call("checkFramebufferStatus", int32(target)).Int())
}

func (R *RenderingContext) GetFramebufferAttachmentParameter(target, attachment, pname glone.Enum) any {
	return castToAny(R.call("getFramebufferAttachmentParameter", int32(target), int32(attachment), int32(pname)))
}

func (R *RenderingContext) BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int32, mask, filter glone.Enum) {
	R.call("blitFramebuffer", srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, int32(mask), int32(filter))
}

func (R *RenderingContext) FramebufferTextureLayer(target, attachment glone.Enum, texture glone.Texture, level, layer int32) {
	t := textureOrNil(texture)
	R.call("framebufferTextureLayer", int32(target), int32(attachment), t, level, layer)
}

func (R *RenderingContext) InvalidateFramebuffer(target glone.Enum, attachments []glone.Enum) {
	R.call("invalidateFramebuffer", int32(target), makeEnumArray(attachments))
}

func (R *RenderingContext) InvalidateSubFramebuffer(target glone.Enum, attachments []glone.Enum, x, y, width, height int32) {
	R.call("invalidateSubFramebuffer", int32(target), makeEnumArray(attachments), x, y, width, height)
}

func (R *RenderingContext) RenderbufferStorage(target, internalformat glone.Enum, width, height int32) {
	R.call("renderbufferStorage", int32(target), int32(internalformat), width, height)
}

func (R *RenderingContext) RenderbufferStorageMultisample(target glone.Enum, samples int32, internalformat glone.Enum, width, height int32) {
	R.call("renderbufferStorageMultisample", int32(target), samples, int32(internalformat), width, height)
}

func (R *RenderingContext) GetRenderbufferParameter(target, pname glone.Enum) any {
	return castToAny(R.call("getRenderbufferParameter", int32(target), int32(pname)))
}

func (R *RenderingContext) GetInternalFormatParameter(target, internalformat, pname glone.Enum) any {
	return castToAny(R.call("getInternalFormatParameter", int32(target), int32(internalformat), int32(pname)))
}

func (R *RenderingContext) ReadPixelsPbo(x, y, width, height int32, format, typ glone.Enum, offset int32) {
	R.call("readPixels", x, y, width, height, int32(format), int32(typ), offset)
}

func (R *RenderingContext) ReadPixelsPix(x, y, width, height int32, format, typ glone.Enum, dstData []byte) {
	dst := makeByteArrayLength(len(dstData))
	R.call("readPixels", x, y, width, height, int32(format), int32(typ), dst)
	js.CopyBytesToGo(dstData, dst)
}

func (R *RenderingContext) DrawBuffers(buffers []glone.Enum) {
	R.call("drawBuffers", makeEnumArray(buffers))
}

func (R *RenderingContext) ClearBufferfv(buffer glone.Enum, drawbuffer int32, values []float32) {
	R.call("clearBufferfv", int32(buffer), drawbuffer, makeFloat32Array(values))
}

func (R *RenderingContext) ClearBufferiv(buffer glone.Enum, drawbuffer int32, values []int32) {
	R.call("clearBufferiv", int32(buffer), drawbuffer, makeInt32Array(values))
}

func (R *RenderingContext) ClearBufferuiv(buffer glone.Enum, drawbuffer int32, values []uint32) {
	R.call("clearBufferuiv", int32(buffer), drawbuffer, makeUint32Array(values))
}

func (R *RenderingContext) ClearBufferfi(buffer glone.Enum, drawbuffer int32, depth float32, stencil int32) {
	R.call("clearBufferfi", int32(buffer), drawbuffer, depth, stencil)
}

func (R *RenderingContext) BufferDataSize(target glone.Enum, size int32, usage glone.Enum) {
	R.call("bufferData", int32(target), size, int32(usage))
}

func (R *RenderingContext) BufferDataSrc(target glone.Enum, data glone.BufferSource, usage glone.Enum) {
	d := bufferSourceOrNil(data)
	R.call("bufferData", int32(target), d, int32(usage))
}

func (R *RenderingContext) BufferDataPix(target glone.Enum, srcData []byte, usage glone.Enum) {
	R.call("bufferData", int32(target), makeByteArray(srcData), int32(usage))
}

func (R *RenderingContext) BufferSubDataSrc(target glone.Enum, offset int32, data glone.BufferSource) {
	d := bufferSourceOrNil(data)
	R.call("bufferSubData", int32(target), offset, d)
}

func (R *RenderingContext) BufferSubDataPix(target glone.Enum, dstByteOffset int32, srcData []byte) {
	R.call("bufferSubData", int32(target), dstByteOffset, makeByteArray(srcData))
}

func (R *RenderingContext) GetBufferParameter(target, pname glone.Enum) any {
	return castToAny(R.call("getBufferParameter", int32(target), int32(pname)))
}

func (R *RenderingContext) CopyBufferSubData(readTarget, writeTarget glone.Enum, readOffset, writeOffset, size int32) {
	R.call("copyBufferSubData", int32(readTarget), int32(writeTarget), readOffset, writeOffset, size)
}

func (R *RenderingContext) GetBufferSubData(target glone.Enum, srcByteOffset int32, dstBuffer []byte) {
	dst := makeByteArrayLength(len(dstBuffer))
	R.call("getBufferSubData", int32(target), srcByteOffset, dst)
	js.CopyBytesToGo(dstBuffer, dst)
}

func (R *RenderingContext) ActiveTexture(texture glone.Enum) {
	R.call("activeTexture", int32(texture))
}

func (R *RenderingContext) TexParameterf(target, pname glone.Enum, param float32) {
	R.call("texParameterf", int32(target), int32(pname), param)
}

func (R *RenderingContext) TexParameteri(target, pname glone.Enum, param int32) {
	R.call("texParameteri", int32(target), int32(pname), param)
}

func (R *RenderingContext) GetTexParameter(target, pname glone.Enum) any {
	return castToAny(R.call("getTexParameter", int32(target), int32(pname)))
}

func (R *RenderingContext) GenerateMipmap(target glone.Enum) {
	R.call("generateMipmap", int32(target))
}

func (R *RenderingContext) CopyTexImage2D(target glone.Enum, level int32, internalformat glone.Enum, x, y, width, height, border int32) {
	R.call("copyTexImage2D", int32(target), level, int32(internalformat), x, y, width, height, border)
}

func (R *RenderingContext) CopyTexSubImage2D(target glone.Enum, level int32, xoffset, yoffset, x, y, width, height int32) {
	R.call("copyTexSubImage2D", int32(target), level, xoffset, yoffset, x, y, width, height)
}

func (R *RenderingContext) TexStorage2D(target glone.Enum, levels int32, internalformat glone.Enum, width, height int32) {
	R.call("texStorage2D", int32(target), levels, int32(internalformat), width, height)
}

func (R *RenderingContext) TexStorage3D(target glone.Enum, levels int32, internalformat glone.Enum, width, height, depth int32) {
	R.call("texStorage3D", int32(target), levels, int32(internalformat), width, height, depth)
}

func (R *RenderingContext) TexImage2DPbo(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, pboOffset int32) {
	R.call("texImage2D", int32(target), level, internalformat, width, height, border, int32(format), int32(typ), pboOffset)
}

func (R *RenderingContext) TexImage2DSrc(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, source glone.TexImageSource) {
	s := texImageSourceOrNil(source)
	R.call("texImage2D", int32(target), level, internalformat, width, height, border, int32(format), int32(typ), s)
}

func (R *RenderingContext) TexImage2DPix(target glone.Enum, level, internalformat, width, height, border int32, format, typ glone.Enum, srcData []byte) {
	R.call("texImage2D", int32(target), level, internalformat, width, height, border, int32(format), int32(typ), makeByteArray(srcData))
}

func (R *RenderingContext) TexSubImage2DPbo(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, pboOffset int32) {
	R.call("texSubImage2D", int32(target), level, xoffset, yoffset, width, height, int32(format), int32(typ), pboOffset)
}

func (R *RenderingContext) TexSubImage2DSrc(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, source glone.TexImageSource) {
	s := texImageSourceOrNil(source)
	R.call("texSubImage2D", int32(target), level, xoffset, yoffset, width, height, int32(format), int32(typ), s)
}

func (R *RenderingContext) TexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int32, format, typ glone.Enum, srcData []byte) {
	R.call("texSubImage2D", int32(target), level, xoffset, yoffset, width, height, int32(format), int32(typ), makeByteArray(srcData))
}

func (R *RenderingContext) TexImage3DPbo(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, pboOffset int32) {
	R.call("texImage3D", int32(target), level, internalformat, width, height, depth, border, int32(format), int32(typ), pboOffset)
}

func (R *RenderingContext) TexImage3DSrc(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, source glone.TexImageSource) {
	s := texImageSourceOrNil(source)
	R.call("texImage3D", int32(target), level, internalformat, width, height, depth, border, int32(format), int32(typ), s)
}

func (R *RenderingContext) TexImage3DPix(target glone.Enum, level, internalformat, width, height, depth, border int32, format, typ glone.Enum, srcData []byte) {
	R.call("texImage3D", int32(target), level, internalformat, width, height, depth, border, int32(format), int32(typ), makeByteArray(srcData))
}

func (R *RenderingContext) TexSubImage3DPbo(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, pboOffset int32) {
	R.call("texSubImage3D", int32(target), level, xoffset, yoffset, zoffset, width, height, depth, int32(format), int32(typ), pboOffset)
}

func (R *RenderingContext) TexSubImage3DSrc(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, source glone.TexImageSource) {
	s := texImageSourceOrNil(source)
	R.call("texSubImage3D", int32(target), level, xoffset, yoffset, zoffset, width, height, depth, int32(format), int32(typ), s)
}

func (R *RenderingContext) TexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ glone.Enum, srcData []byte) {
	R.call("texSubImage3D", int32(target), level, xoffset, yoffset, zoffset, width, height, depth, int32(format), int32(typ), makeByteArray(srcData))
}

func (R *RenderingContext) CopyTexSubImage3D(target glone.Enum, level, xoffset, yoffset, zoffset, x, y, width, height int32) {
	R.call("copyTexSubImage3D", int32(target), level, xoffset, yoffset, zoffset, x, y, width, height)
}

func (R *RenderingContext) CompressedTexImage2DSize(target glone.Enum, level int32, internalformat glone.Enum, width, height, border, imageSize, offset int32) {
	R.call("compressedTexImage2D", int32(target), level, int32(internalformat), width, height, border, imageSize, offset)
}

func (R *RenderingContext) CompressedTexImage2DPix(target glone.Enum, level int32, internalformat glone.Enum, width, height, border int32, srcData []byte) {
	R.call("compressedTexImage2D", int32(target), level, int32(internalformat), width, height, border, makeByteArray(srcData))
}

func (R *RenderingContext) CompressedTexSubImage2DSize(target glone.Enum, level, xoffset, yoffset, width, height int32, format glone.Enum, imageSize, offset int32) {
	R.call("compressedTexSubImage2D", int32(target), level, xoffset, yoffset, width, height, int32(format), imageSize, offset)
}

func (R *RenderingContext) CompressedTexSubImage2DPix(target glone.Enum, level, xoffset, yoffset, width, height int32, format glone.Enum, srcData []byte) {
	R.call("compressedTexSubImage2D", int32(target), level, xoffset, yoffset, width, height, int32(format), makeByteArray(srcData))
}

func (R *RenderingContext) CompressedTexImage3DSize(target glone.Enum, level int32, internalformat glone.Enum, width, height, depth, border, imageSize, offset int32) {
	R.call("compressedTexImage3D", int32(target), level, int32(internalformat), width, height, depth, border, imageSize, offset)
}

func (R *RenderingContext) CompressedTexImage3DPix(target glone.Enum, level int32, internalformat glone.Enum, width, height, depth, border int32, srcData []byte) {
	R.call("compressedTexImage3D", int32(target), level, int32(internalformat), width, height, depth, border, makeByteArray(srcData))
}

func (R *RenderingContext) CompressedTexSubImage3DSize(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format glone.Enum, imageSize, offset int32) {
	R.call("compressedTexSubImage3D", int32(target), level, xoffset, yoffset, zoffset, width, height, depth, int32(format), imageSize, offset)
}

func (R *RenderingContext) CompressedTexSubImage3DPix(target glone.Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format glone.Enum, srcData []byte) {
	R.call("compressedTexSubImage3D", int32(target), level, xoffset, yoffset, zoffset, width, height, depth, int32(format), makeByteArray(srcData))
}

func (R *RenderingContext) BeginQuery(target glone.Enum, query glone.Query) {
	q := queryOrNil(query)
	R.call("beginQuery", int32(target), q)
}

func (R *RenderingContext) EndQuery(target glone.Enum) {
	R.call("endQuery", int32(target))
}

func (R *RenderingContext) GetQuery(target, pname glone.Enum) glone.Query {
	return &Query{
		R.call("getQuery", int32(target), int32(pname)),
	}
}

func (R *RenderingContext) GetQueryParameter(query glone.Query, pname glone.Enum) any {
	q := queryOrNil(query)
	return castToAny(R.call("getQueryParameter", q, int32(pname)))
}

func (R *RenderingContext) SamplerParameteri(sampler glone.Sampler, pname glone.Enum, param int32) {
	s := samplerOrNil(sampler)
	R.call("samplerParameteri", s, int32(pname), param)
}

func (R *RenderingContext) SamplerParameterf(sampler glone.Sampler, pname glone.Enum, param float32) {
	s := samplerOrNil(sampler)
	R.call("samplerParameterf", s, int32(pname), param)
}

func (R *RenderingContext) GetSamplerParameter(sampler glone.Sampler, pname glone.Enum) any {
	s := samplerOrNil(sampler)
	return castToAny(R.call("getSamplerParameter", s, int32(pname)))
}

func (R *RenderingContext) FenceSync(condition, flags glone.Enum) glone.Sync {
	return &Sync{
		R.call("fenceSync", int32(condition), int32(flags)),
	}
}

func (R *RenderingContext) ClientWaitSync(sync glone.Sync, flags glone.Enum, timeout uint64) {
	s := syncOrNil(sync)
	R.call("clientWaitSync", s, int32(flags), timeout)
}

func (R *RenderingContext) WaitSync(sync glone.Sync, flags glone.Enum, timeout int64) {
	s := syncOrNil(sync)
	R.call("waitSync", s, int32(flags), timeout)
}

func (R *RenderingContext) GetSyncParameter(sync glone.Sync, pname glone.Enum) any {
	s := syncOrNil(sync)
	return castToAny(R.call("getSyncParameter", s, int32(pname)))
}

func (R *RenderingContext) BeginTransformFeedback(primitiveMode glone.Enum) {
	R.call("beginTransformFeedback", int32(primitiveMode))
}

func (R *RenderingContext) EndTransformFeedback() {
	R.call("endTransformFeedback")
}

func (R *RenderingContext) TransformFeedbackVaryings(program glone.Program, varyings []string, bufferMode glone.Enum) {
	p := programOrNil(program)
	R.call("transformFeedbackVaryings", p, makeStringArray(varyings), int32(bufferMode))
}

func (R *RenderingContext) GetTransformFeedbackVarying(program glone.Program, index uint32) glone.ActiveInfo {
	p := programOrNil(program)
	return &ActiveInfo{
		R.call("getTransformFeedbackVarying", p, index),
	}
}

func (R *RenderingContext) PauseTransformFeedback() {
	R.call("pauseTransformFeedback")
}

func (R *RenderingContext) ResumeTransformFeedback() {
	R.call("resumeTransformFeedback")
}

func (R *RenderingContext) GetIndexedParameter(target glone.Enum, index uint32) any {
	return castToAny(R.call("getIndexedParameter", int32(target), index))
}

func (R *RenderingContext) GetUniformIndices(program glone.Program, uniformNames []string) []uint32 {
	p := programOrNil(program)
	return getUintArray(R.call("getUniformIndices", p, makeStringArray(uniformNames)))
}

func (R *RenderingContext) GetActiveUniforms(program glone.Program, uniformIndices []uint32, pname glone.Enum) any {
	p := programOrNil(program)
	return castToAny(R.call("getActiveUniforms", p, makeUintArray(uniformIndices), int32(pname)))
}

func (R *RenderingContext) GetUniformBlockIndex(program glone.Program, uniformBlockName string) uint32 {
	p := programOrNil(program)
	return uint32(R.call("getUniformBlockIndex", p, uniformBlockName).Int())
}

func (R *RenderingContext) GetActiveUniformBlockParameter(program glone.Program, uniformBlockIndex uint32, pname glone.Enum) any {
	p := programOrNil(program)
	return castToAny(R.call("getActiveUniformBlockParameter", p, uniformBlockIndex, int32(pname)))
}

func (R *RenderingContext) GetActiveUniformBlockName(program glone.Program, uniformBlockIndex uint32) string {
	p := programOrNil(program)
	return R.call("getActiveUniformBlockName", p, uniformBlockIndex).String()
}

func (R *RenderingContext) UniformBlockBinding(program glone.Program, uniformBlockIndex, uniformBlockBinding uint32) {
	p := programOrNil(program)
	R.call("uniformBlockBinding", p, uniformBlockIndex, uniformBlockBinding)
}

func (R *RenderingContext) Viewport(x, y, width, height int32) {
	R.call("viewport", x, y, width, height)
}

func (R *RenderingContext) Scissor(x, y, width, height int32) {
	R.call("scissor", x, y, width, height)
}

func (R *RenderingContext) BlendColor(red, green, blue, alpha float32) {
	R.call("blendColor", red, green, blue, alpha)
}

func (R *RenderingContext) BlendEquation(mode glone.Enum) {
	R.call("blendEquation", int32(mode))
}

func (R *RenderingContext) BlendEquationSeparate(modeRGB, modeAlpha glone.Enum) {
	R.call("blendEquationSeparate", int32(modeRGB), int32(modeAlpha))
}

func (R *RenderingContext) BlendFunc(sfactor, dfactor glone.Enum) {
	R.call("blendFunc", int32(sfactor), int32(dfactor))
}

func (R *RenderingContext) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha glone.Enum) {
	R.call("blendFuncSeparate", int32(srcRGB), int32(dstRGB), int32(srcAlpha), int32(dstAlpha))
}

func (R *RenderingContext) Clear(mask glone.Enum) {
	R.call("clear", int32(mask))
}

func (R *RenderingContext) ClearColor(red, green, blue, alpha float32) {
	R.call("clearColor", red, green, blue, alpha)
}

func (R *RenderingContext) ClearDepth(depth float32) {
	R.call("clearDepth", depth)
}

func (R *RenderingContext) ClearStencil(s int32) {
	R.call("clearStencil", s)
}

func (R *RenderingContext) ColorMask(red, green, blue, alpha bool) {
	R.call("colorMask", red, green, blue, alpha)
}

func (R *RenderingContext) CullFace(mode glone.Enum) {
	R.call("cullFace", int32(mode))
}

func (R *RenderingContext) FrontFace(mode glone.Enum) {
	R.call("frontFace", int32(mode))
}

func (R *RenderingContext) DepthFunc(fun glone.Enum) {
	R.call("depthFunc", int32(fun))
}

func (R *RenderingContext) DepthMask(flag bool) {
	R.call("depthMask", flag)
}

func (R *RenderingContext) DepthRange(zNear, zFar float32) {
	R.call("depthRange", zNear, zFar)
}

func (R *RenderingContext) Disable(cap glone.Enum) {
	R.call("disable", int32(cap))
}

func (R *RenderingContext) Enable(cap glone.Enum) {
	R.call("enable", int32(cap))
}

func (R *RenderingContext) Hint(target, mode glone.Enum) {
	R.call("hint", int32(target), int32(mode))
}

func (R *RenderingContext) LineWidth(width float32) {
	R.call("lineWidth", width)
}

func (R *RenderingContext) PixelStorei(pname glone.Enum, param int32) {
	R.call("pixelStorei", int32(pname), param)
}

func (R *RenderingContext) PolygonOffset(factor, units float32) {
	R.call("polygonOffset", factor, units)
}

func (R *RenderingContext) StencilFunc(fun glone.Enum, ref int32, mask uint32) {
	R.call("stencilFunc", int32(fun), ref, mask)
}

func (R *RenderingContext) StencilFuncSeparate(face, fun glone.Enum, ref int32, mask uint32) {
	R.call("stencilFuncSeparate", int32(face), int32(fun), ref, mask)
}

func (R *RenderingContext) StencilMask(mask uint32) {
	R.call("stencilMask", mask)
}

func (R *RenderingContext) StencilMaskSeparate(face glone.Enum, mask uint32) {
	R.call("stencilMaskSeparate", int32(face), mask)
}

func (R *RenderingContext) StencilOp(fail, zfail, zpass glone.Enum) {
	R.call("stencilOp", int32(fail), int32(zfail), int32(zpass))
}

func (R *RenderingContext) StencilOpSeparate(face, fail, zfail, zpass glone.Enum) {
	R.call("stencilOpSeparate", int32(face), int32(fail), int32(zfail), int32(zpass))
}

func (R *RenderingContext) SampleCoverage(value float32, invert bool) {
	R.call("sampleCoverage", value, invert)
}

func (R *RenderingContext) ReadBuffer(src glone.Enum) {
	R.call("readBuffer", int32(src))
}

func (R *RenderingContext) GetParameter(pname glone.Enum) any {
	return castToAny(R.call("getParameter", int32(pname)))
}

func (R *RenderingContext) Finish() {
	R.call("finish")
}

func (R *RenderingContext) Flush() {
	R.call("flush")
}

func (R *RenderingContext) DrawArrays(mode glone.Enum, first, count int32) {
	R.call("drawArrays", int32(mode), first, count)
}

func (R *RenderingContext) DrawElements(mode glone.Enum, count int32, typ glone.Enum, offset int32) {
	R.call("drawElements", int32(mode), count, int32(typ), offset)
}

func (R *RenderingContext) DrawArraysInstanced(mode glone.Enum, first, count, instanceCount int32) {
	R.call("drawArraysInstanced", int32(mode), first, count, instanceCount)
}

func (R *RenderingContext) DrawElementsInstanced(mode glone.Enum, count int32, typ glone.Enum, offset, instanceCount int32) {
	R.call("drawElementsInstanced", int32(mode), count, int32(typ), offset, instanceCount)
}

func (R *RenderingContext) DrawRangeElements(mode glone.Enum, start, end uint32, count int32, typ glone.Enum, offset int32) {
	R.call("drawRangeElements", int32(mode), start, end, count, int32(typ), offset)
}

func (R *RenderingContext) GetError() glone.Enum {
	return glone.Enum(R.call("getError").Int())
}

var _ glone.RenderingContext = (*RenderingContext)(nil)
