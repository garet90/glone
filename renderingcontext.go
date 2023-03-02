package glone

type Program interface {
	GLOneProgram()
}
type Shader interface {
	GLOneShader()
}
type Buffer interface {
	GLOneBuffer()
}
type Framebuffer interface {
	GLOneFramebuffer()
}
type Renderbuffer interface {
	GLOneRenderbuffer()
}
type Texture interface {
	GLOneTexture()
}
type ActiveInfo interface {
	GLOneActiveInfo()
}
type ShaderPrecisionFormat interface {
	GLOneShaderPrecisionFormat()
}
type UniformLocation interface {
	GLOneUniformLocation()
}
type BufferSource interface {
	GLOneBufferSource()
}
type TexImageSource interface {
	GLOneTexImageSource()
}
type Query interface {
	GLOneQuery()
}
type Sampler interface {
	GLOneSampler()
}
type Sync interface {
	GLOneSync()
}
type TransformFeedback interface {
	GLOneTransformFeedback()
}
type VertexArray interface {
	GLOneVertexArray()
}

type RenderingContext interface {
	/* Object Creation */

	CreateBuffer() Buffer
	CreateFramebuffer() Framebuffer
	CreateProgram() Program
	CreateRenderbuffer() Renderbuffer
	CreateShader(typ Enum) Shader
	CreateTexture() Texture
	CreateVertexArray() VertexArray
	CreateTransformFeedback() TransformFeedback
	CreateSampler() Sampler
	CreateQuery() Query

	/* Object Binding */

	BindAttribLocation(program Program, index uint32, name string)
	BindBuffer(target Enum, buffer Buffer)
	BindBufferBase(target Enum, index uint32, buffer Buffer)
	BindBufferRange(target Enum, index uint32, buffer Buffer, offset, size int32)
	BindFramebuffer(target Enum, framebuffer Framebuffer)
	BindRenderbuffer(target Enum, renderbuffer Renderbuffer)
	BindTexture(target Enum, texture Texture)
	BindVertexArray(array VertexArray)
	BindTransformFeedback(target Enum, tf TransformFeedback)
	BindSampler(unit uint32, sampler Sampler)

	/* Object Type Checking */

	IsBuffer(buffer Buffer) bool
	IsEnabled(cap Enum) bool
	IsFramebuffer(framebuffer Framebuffer) bool
	IsProgram(program Program) bool
	IsRenderbuffer(renderbuffer Renderbuffer) bool
	IsShader(shader Shader) bool
	IsTexture(texture Texture) bool
	IsVertexArray(vertexArray VertexArray) bool
	IsTransformFeedback(tf TransformFeedback) bool
	IsSampler(sampler Sampler) bool
	IsQuery(query Query) bool
	IsSync(sync Sync) bool

	/* Object Deletion */

	DeleteBuffer(buffer Buffer)
	DeleteFramebuffer(framebuffer Framebuffer)
	DeleteProgram(program Program)
	DeleteRenderbuffer(renderbuffer Renderbuffer)
	DeleteShader(shader Shader)
	DeleteTexture(texture Texture)
	DeleteVertexArray(vertexArray VertexArray)
	DeleteTransformFeedback(tf TransformFeedback)
	DeleteSampler(sampler Sampler)
	DeleteQuery(query Query)
	DeleteSync(sync Sync)

	/* Shader Methods */

	ShaderSource(shader Shader, source string)
	CompileShader(shader Shader)
	GetShaderSource(shader Shader) string
	GetShaderParameter(shader Shader, pname Enum) any
	GetShaderPrecisionFormat(shadertype, precisiontype Enum) ShaderPrecisionFormat
	GetShaderInfoLog(shader Shader) string

	/* Program Methods */

	AttachShader(program Program, shader Shader)
	DetachShader(program Program, shader Shader)
	LinkProgram(program Program)
	ValidateProgram(program Program)
	UseProgram(program Program)
	GetActiveAttrib(program Program, index uint32) ActiveInfo
	GetActiveUniform(program Program, index uint32) ActiveInfo
	GetAttachedShaders(program Program) []Shader
	GetAttribLocation(program Program, name string) int32
	GetProgramInfoLog(program Program) string
	GetUniform(program Program, location UniformLocation) any
	GetUniformLocation(program Program, name string) UniformLocation
	GetFragDataLocation(program Program, name string) int32
	GetProgramParameter(program Program, pname Enum) any
	TransformFeedbackVaryings(program Program, varyings []string, bufferMode Enum)
	GetTransformFeedbackVarying(program Program, index uint32) ActiveInfo

	/* Program Uniform Bindings */

	Uniform1f(location UniformLocation, x float32)
	Uniform2f(location UniformLocation, x, y float32)
	Uniform3f(location UniformLocation, x, y, z float32)
	Uniform4f(location UniformLocation, x, y, z, w float32)

	Uniform1i(location UniformLocation, x int32)
	Uniform2i(location UniformLocation, x, y int32)
	Uniform3i(location UniformLocation, x, y, z int32)
	Uniform4i(location UniformLocation, x, y, z, w int32)

	Uniform1ui(location UniformLocation, v0 uint32)
	Uniform2ui(location UniformLocation, v0, v1 uint32)
	Uniform3ui(location UniformLocation, v0, v1, v2 uint32)
	Uniform4ui(location UniformLocation, v0, v1, v2, v3 uint32)

	Uniform1fv(location UniformLocation, v []float32)
	Uniform2fv(location UniformLocation, v []float32)
	Uniform3fv(location UniformLocation, v []float32)
	Uniform4fv(location UniformLocation, v []float32)

	Uniform1iv(location UniformLocation, v []int32)
	Uniform2iv(location UniformLocation, v []int32)
	Uniform3iv(location UniformLocation, v []int32)
	Uniform4iv(location UniformLocation, v []int32)

	Uniform1uiv(location UniformLocation, data []uint32)
	Uniform2uiv(location UniformLocation, data []uint32)
	Uniform3uiv(location UniformLocation, data []uint32)
	Uniform4uiv(location UniformLocation, data []uint32)

	UniformMatrix2fv(location UniformLocation, transpose bool, value []float32)
	UniformMatrix3fv(location UniformLocation, transpose bool, value []float32)
	UniformMatrix4fv(location UniformLocation, transpose bool, value []float32)

	UniformMatrix3x2fv(location UniformLocation, transpose bool, data []float32)
	UniformMatrix4x2fv(location UniformLocation, transpose bool, data []float32)

	UniformMatrix2x3fv(location UniformLocation, transpose bool, data []float32)
	UniformMatrix4x3fv(location UniformLocation, transpose bool, data []float32)

	UniformMatrix2x4fv(location UniformLocation, transpose bool, data []float32)
	UniformMatrix3x4fv(location UniformLocation, transpose bool, data []float32)

	/* Vertex Attribute Object Bindings */

	EnableVertexAttribArray(index uint32)
	DisableVertexAttribArray(index uint32)

	VertexAttrib1f(index uint32, x float32)
	VertexAttrib2f(index uint32, x, y float32)
	VertexAttrib3f(index uint32, x, y, z float32)
	VertexAttrib4f(index uint32, x, y, z, w float32)

	VertexAttrib1fv(index uint32, values []float32)
	VertexAttrib2fv(index uint32, values []float32)
	VertexAttrib3fv(index uint32, values []float32)
	VertexAttrib4fv(index uint32, values []float32)

	VertexAttribI4i(index uint32, x, y, z, w int32)
	VertexAttribI4ui(index, x, y, z, w uint32)

	VertexAttribI4iv(index uint32, values []int32)
	VertexAttribI4uiv(index uint32, values []uint32)

	VertexAttribPointer(index uint32, size int32, typ Enum, normalized bool, stride, offset int32)
	VertexAttribIPointer(index uint32, size int32, typ Enum, stride, offset int32)

	VertexAttribDivisor(index, divisor uint32)

	GetVertexAttrib(index uint32, pname Enum) any
	GetVertexAttribOffset(index uint32, pname Enum) int32

	/* Framebuffer Methods */

	FramebufferRenderbuffer(target, attachment, renderbuffertarget Enum, renderbuffer Renderbuffer)
	FramebufferTexture2D(target, attachment, textarget Enum, texture Texture, level int32)
	CheckFramebufferStatus(target Enum) Enum
	GetFramebufferAttachmentParameter(target, attachment, pname Enum) any
	FramebufferTextureLayer(target, attachment Enum, texture Texture, level, layer int32)
	InvalidateFramebuffer(target Enum, attachments []Enum)
	InvalidateSubFramebuffer(target Enum, attachments []Enum, x, y, width, height int32)
	BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int32, mask, filter Enum)

	/* Renderbuffer Methods */

	RenderbufferStorage(target, internalformat Enum, width, height int32)
	RenderbufferStorageMultisample(target Enum, samples int32, internalformat Enum, width, height int32)
	GetRenderbufferParameter(target, pname Enum) any
	GetInternalFormatParameter(target, internalformat, pname Enum) any

	/* Pixel Feedback Methods */

	ReadPixelsPbo(x, y, width, height int32, format, typ Enum, offset int32)
	ReadPixelsPix(x, y, width, height int32, format, typ Enum, dstData []byte)

	/* Buffer Methods */

	BufferDataSize(target Enum, size int32, usage Enum)
	BufferDataSrc(target Enum, data BufferSource, usage Enum)
	BufferDataPix(target Enum, srcData []byte, usage Enum)
	BufferSubDataSrc(target Enum, offset int32, data BufferSource)
	BufferSubDataPix(target Enum, dstByteOffset int32, srcData []byte)
	GetBufferParameter(target, pname Enum) any
	CopyBufferSubData(readTarget, writeTarget Enum, readOffset, writeOffset, size int32)
	GetBufferSubData(target Enum, srcByteOffset int32, dstBuffer []byte)

	/* Texture Methods */

	ActiveTexture(texture Enum)

	TexParameterf(target, pname Enum, param float32)
	TexParameteri(target, pname Enum, param int32)
	GetTexParameter(target, pname Enum) any
	GenerateMipmap(target Enum)

	CopyTexImage2D(target Enum, level int32, internalformat Enum, x, y, width, height, border int32)
	CopyTexSubImage2D(target Enum, level int32, xoffset, yoffset, x, y, width, height int32)
	CopyTexSubImage3D(target Enum, level, xoffset, yoffset, zoffset, x, y, width, height int32)

	TexStorage2D(target Enum, levels int32, internalformat Enum, width, height int32)
	TexStorage3D(target Enum, levels int32, internalformat Enum, width, height, depth int32)

	TexImage2DPbo(target Enum, level, internalformat, width, height, border int32, format, typ Enum, pboOffset int32)
	TexImage2DSrc(target Enum, level, internalformat, width, height, border int32, format, typ Enum, source TexImageSource)
	TexImage2DPix(target Enum, level, internalformat, width, height, border int32, format, typ Enum, srcData []byte)

	TexSubImage2DPbo(target Enum, level, xoffset, yoffset, width, height int32, format, typ Enum, pboOffset int32)
	TexSubImage2DSrc(target Enum, level, xoffset, yoffset, width, height int32, format, typ Enum, source TexImageSource)
	TexSubImage2DPix(target Enum, level, xoffset, yoffset, width, height int32, format, typ Enum, srcData []byte)

	TexImage3DPbo(target Enum, level, internalformat, width, height, depth, border int32, format, typ Enum, pboOffset int32)
	TexImage3DSrc(target Enum, level, internalformat, width, height, depth, border int32, format, typ Enum, source TexImageSource)
	TexImage3DPix(target Enum, level, internalformat, width, height, depth, border int32, format, typ Enum, srcData []byte)

	TexSubImage3DPbo(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ Enum, pboOffset int32)
	TexSubImage3DSrc(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ Enum, source TexImageSource)
	TexSubImage3DPix(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format, typ Enum, srcData []byte)

	CompressedTexImage2DSize(target Enum, level int32, internalformat Enum, width, height, border, imageSize, offset int32)
	CompressedTexImage2DPix(target Enum, level int32, internalformat Enum, width, height, border int32, srcData []byte)

	CompressedTexSubImage2DSize(target Enum, level, xoffset, yoffset, width, height int32, format Enum, imageSize, offset int32)
	CompressedTexSubImage2DPix(target Enum, level, xoffset, yoffset, width, height int32, format Enum, srcData []byte)

	CompressedTexImage3DSize(target Enum, level int32, internalformat Enum, width, height, depth, border, imageSize, offset int32)
	CompressedTexImage3DPix(target Enum, level int32, internalformat Enum, width, height, depth, border int32, srcData []byte)

	CompressedTexSubImage3DSize(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format Enum, imageSize, offset int32)
	CompressedTexSubImage3DPix(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int32, format Enum, srcData []byte)

	/* Query Methods */

	BeginQuery(target Enum, query Query)
	EndQuery(target Enum)
	GetQuery(target, pname Enum) Query
	GetQueryParameter(query Query, pname Enum) any

	/* Sampler Methods */

	SamplerParameteri(sampler Sampler, pname Enum, param int32)
	SamplerParameterf(sampler Sampler, pname Enum, param float32)
	GetSamplerParameter(sampler Sampler, pname Enum) any

	/* Sync Methods */

	FenceSync(condition, flags Enum) Sync
	ClientWaitSync(sync Sync, flags Enum, timeout uint64)
	WaitSync(sync Sync, flags Enum, timeout int64)
	GetSyncParameter(sync Sync, pname Enum) any

	/* Transform Feedback Methods */

	BeginTransformFeedback(primitiveMode Enum)
	EndTransformFeedback()
	PauseTransformFeedback()
	ResumeTransformFeedback()

	/* Transform Feedback / Uniform Buffer Object Methods */

	GetIndexedParameter(target Enum, index uint32) any
	GetUniformIndices(program Program, uniformNames []string) []uint32
	GetActiveUniforms(program Program, uniformIndices []uint32, pname Enum) any
	GetUniformBlockIndex(program Program, uniformBlockName string) uint32
	GetActiveUniformBlockParameter(program Program, uniformBlockIndex uint32, pname Enum) any
	GetActiveUniformBlockName(program Program, uniformBlockIndex uint32) string
	UniformBlockBinding(program Program, uniformBlockIndex, uniformBlockBinding uint32)

	/* Viewport Methods */

	Viewport(x, y, width, height int32)
	Scissor(x, y, width, height int32)

	/* Clear Methods */

	Clear(mask Enum)
	ClearBufferfv(buffer Enum, drawbuffer int32, values []float32)
	ClearBufferiv(buffer Enum, drawbuffer int32, values []int32)
	ClearBufferuiv(buffer Enum, drawbuffer int32, values []uint32)
	ClearBufferfi(buffer Enum, drawbuffer int32, depth float32, stencil int32)

	/* State Methods */

	BlendColor(red, green, blue, alpha float32)
	BlendEquation(mode Enum)
	BlendEquationSeparate(modeRGB, modeAlpha Enum)
	BlendFunc(sfactor, dfactor Enum)
	BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha Enum)
	ClearColor(red, green, blue, alpha float32)
	ClearDepth(depth float32)
	ClearStencil(s int32)
	ColorMask(red, green, blue, alpha bool)
	CullFace(mode Enum)
	FrontFace(mode Enum)
	DepthFunc(fun Enum)
	DepthMask(flag bool)
	DepthRange(zNear, zFar float32)
	Disable(cap Enum)
	Enable(cap Enum)
	Hint(target, mode Enum)
	LineWidth(width float32)
	PixelStorei(pname Enum, param int32)
	PolygonOffset(factor, units float32)
	StencilFunc(fun Enum, ref int32, mask uint32)
	StencilFuncSeparate(face, fun Enum, ref int32, mask uint32)
	StencilMask(mask uint32)
	StencilMaskSeparate(face Enum, mask uint32)
	StencilOp(fail, zfail, zpass Enum)
	StencilOpSeparate(face, fail, zfail, zpass Enum)
	SampleCoverage(value float32, invert bool)
	ReadBuffer(src Enum)
	DrawBuffers(buffers []Enum)

	GetParameter(pname Enum) any

	/* Pipeline Control Methods */

	Finish()
	Flush()

	/* Draw Methods */

	DrawArrays(mode Enum, first, count int32)
	DrawElements(mode Enum, count int32, typ Enum, offset int32)
	DrawArraysInstanced(mode Enum, first, count, instanceCount int32)
	DrawElementsInstanced(mode Enum, count int32, typ Enum, offset, instanceCount int32)
	DrawRangeElements(mode Enum, start, end uint32, count int32, typ Enum, offset int32)

	/* Error Methods */

	GetError() Enum
}
