package glone

type Program interface{}
type Shader interface{}
type Buffer interface{}
type Framebuffer interface{}
type Renderbuffer interface{}
type Texture interface{}
type ActiveInfo interface{}
type ShaderPrecisionFormat interface{}
type UniformLocation interface{}
type BufferSource interface{}
type TexImageSource interface{}
type Query interface{}
type Sampler interface{}
type Sync interface{}
type TransformFeedback interface{}
type VertexArray interface{}

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

	BindAttribLocation(program Program, index uint, name string)
	BindBuffer(target Enum, buffer Buffer)
	BindBufferBase(target Enum, index uint, buffer Buffer)
	BindBufferRange(target Enum, index uint, buffer Buffer, offset, size int)
	BindFramebuffer(target Enum, framebuffer Framebuffer)
	BindRenderbuffer(target Enum, renderbuffer Renderbuffer)
	BindTexture(target Enum, texture Texture)
	BindVertexArray(array VertexArray)
	BindTransformFeedback(target Enum, tf TransformFeedback)
	BindSampler(unit uint, sampler Sampler)

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
	GetActiveAttrib(program Program, index uint) ActiveInfo
	GetActiveUniform(program Program, index uint) ActiveInfo
	GetAttachedShaders(program Program) []Shader
	GetAttribLocation(program Program, name string) int
	GetProgramInfoLog(program Program) string
	GetUniform(program Program, location UniformLocation) any
	GetUniformLocation(program Program, name string) UniformLocation
	GetFragDataLocation(program Program, name string) int
	GetProgramParameter(program Program, pname Enum) any

	/* Program Uniform Bindings */

	Uniform1f(location UniformLocation, x float64)
	Uniform2f(location UniformLocation, x, y float64)
	Uniform3f(location UniformLocation, x, y, z float64)
	Uniform4f(location UniformLocation, x, y, z, w float64)

	Uniform1i(location UniformLocation, x int)
	Uniform2i(location UniformLocation, x, y int)
	Uniform3i(location UniformLocation, x, y, z int)
	Uniform4i(location UniformLocation, x, y, z, w int)

	Uniform1ui(location UniformLocation, v0 uint)
	Uniform2ui(location UniformLocation, v0, v1 uint)
	Uniform3ui(location UniformLocation, v0, v1, v2 uint)
	Uniform4ui(location UniformLocation, v0, v1, v2, v3 uint)

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

	EnableVertexAttribArray(index uint)
	DisableVertexAttribArray(index uint)

	VertexAttrib1f(index uint, x float64)
	VertexAttrib2f(index uint, x, y float64)
	VertexAttrib3f(index uint, x, y, z float64)
	VertexAttrib4f(index uint, x, y, z, w float64)

	VertexAttrib1fv(index uint, values []float32)
	VertexAttrib2fv(index uint, values []float32)
	VertexAttrib3fv(index uint, values []float32)
	VertexAttrib4fv(index uint, values []float32)

	VertexAttribI4i(index uint, x, y, z, w int)
	VertexAttribI4ui(index, x, y, z, w uint)

	VertexAttribI4iv(index uint, values []int32)
	VertexAttribI4uiv(index uint, values []uint32)

	VertexAttribPointer(index uint, size int, typ Enum, normalized bool, stride, offset int)
	VertexAttribIPointer(index uint, size int, typ Enum, stride, offset int)

	VertexAttribDivisor(index, divisor uint)

	GetVertexAttrib(index uint, pname Enum) any
	GetVertexAttribOffset(index uint, pname Enum) int

	/* Framebuffer Methods */

	FramebufferRenderbuffer(target, attachment, renderbuffertarget Enum, renderbuffer Renderbuffer)
	FramebufferTexture2D(target, attachment, textarget Enum, texture Texture, level int)
	CheckFramebufferStatus(target Enum) Enum
	GetFramebufferAttachmentParameter(target, attachment, pname Enum) any
	BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask, filter Enum)
	FramebufferTextureLayer(target, attachment Enum, texture Texture, level, layer int)
	InvalidateFramebuffer(target Enum, attachments []Enum)
	InvalidateSubFramebuffer(target Enum, attachments []Enum, x, y, width, height int)
	ReadBuffer(src Enum)

	/* Renderbuffer Methods */

	RenderbufferStorage(target, internalformat Enum, width, height int)
	RenderbufferStorageMultisample(target Enum, samples int, internalformat Enum, width, height int)
	GetRenderbufferParameter(target, pname Enum) any
	GetInternalFormatParameter(target, internalformat, pname Enum) any

	/* Pixel Feedback Methods */

	ReadPixelsOff(x, y, width, height int, format, typ Enum, offset int)
	ReadPixelsPix(x, y, width, height int, format, typ Enum, dstData []byte)

	/* Multiple Draw Buffer Methods */

	DrawBuffers(buffers []Enum)
	ClearBufferfv(buffer Enum, drawbuffer int, values []float32)
	ClearBufferiv(buffer Enum, drawbuffer int, values []int32)
	ClearBufferuiv(buffer Enum, drawbuffer int, values []uint32)
	ClearBufferfi(buffer Enum, drawbuffer int, depth float64, stencil int)

	/* Buffer Methods */

	BufferDataSize(target Enum, size int, usage Enum)
	BufferDataSrc(target Enum, data BufferSource, usage Enum)
	BufferDataPix(target Enum, srcData []byte, usage Enum)
	BufferSubDataSrc(target Enum, offset int, data BufferSource)
	BufferSubDataPix(target Enum, dstByteOffset int, srcData []byte)
	GetBufferParameter(target, pname Enum) any
	CopyBufferSubData(readTarget, writeTarget Enum, readOffset, writeOffset, size int)
	GetBufferSubData(target Enum, srcByteOffset int, dstBuffer []byte)

	/* Texture Methods */

	ActiveTexture(texture Enum)

	TexParameterf(target, pname Enum, param float64)
	TexParameteri(target, pname Enum, param int)
	GetTexParameter(target, pname Enum) any
	GenerateMipmap(target Enum)

	CopyTexImage2D(target Enum, level int, internalformat Enum, x, y, width, height, border int)
	CopyTexSubImage2D(target Enum, level int, xoffset, yoffset, x, y, width, height int)

	TexStorage2D(target Enum, levels int, internalformat Enum, width, height int)
	TexStorage3D(target Enum, levels int, internalformat Enum, width, height, depth int)

	TexImage2DPbo(target Enum, level, internalformat, width, height, border int, format, typ Enum, pboOffset int)
	TexImage2DSrc(target Enum, level, internalformat, width, height, border int, format, typ Enum, source TexImageSource)
	TexImage2DPix(target Enum, level, internalformat, width, height, border int, format, typ Enum, srcData []byte)

	TexSubImage2DPbo(target Enum, level, xoffset, yoffset, width, height int, format, typ Enum, pboOffset int)
	TexSubImage2DSrc(target Enum, level, xoffset, yoffset, width, height int, format, typ Enum, source TexImageSource)
	TexSubImage2DPix(target Enum, level, xoffset, yoffset, width, height int, format, typ Enum, srcData []byte)

	TexImage3DPbo(target Enum, level, internalformat, width, height, depth, border int, format, typ Enum, pboOffset int)
	TexImage3DSrc(target Enum, level, internalformat, width, height, depth, border int, format, typ Enum, source TexImageSource)
	TexImage3DPix(target Enum, level, internalformat, width, height, depth, border int, format, typ Enum, srcData []byte)

	TexSubImage3DPbo(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ Enum, pboOffset int)
	TexSubImage3DSrc(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ Enum, source TexImageSource)
	TexSubImage3DPix(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format, typ Enum, srcData []byte)

	CopyTexSubImage3D(target Enum, level, xoffset, yoffset, zoffset, x, y, width, height int)

	CompressedTexImage2DOff(target Enum, level int, internalformat Enum, width, height, border, imageSize, offset int)
	CompressedTexImage2DPix(target Enum, level int, internalformat Enum, width, height, border int, srcData []byte)

	CompressedTexSubImage2DOff(target Enum, level, xoffset, yoffset, width, height int, format Enum, imageSize, offset int)
	CompressedTexSubImage2DPix(target Enum, level, xoffset, yoffset, width, height int, format Enum, srcData []byte)

	CompressedTexImage3D(target Enum, level int, internalformat Enum, width, height, depth, border, imageSize, offset int)
	CompressedTexImage3DSrc(target Enum, level int, internalformat Enum, width, height, depth, border int, srcData []byte)

	CompressedTexSubImage3D(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format Enum, imageSize, offset int)
	CompressedTexSubImage3DSrc(target Enum, level, xoffset, yoffset, zoffset, width, height, depth int, format Enum, srcData []byte)

	/* Query Methods */

	BeginQuery(target Enum, query Query)
	EndQuery(target Enum)
	GetQuery(target, pname Enum) Query
	GetQueryParameter(query Query, pname Enum) any

	/* Sampler Methods */

	SamplerParameteri(sampler Sampler, pname Enum, param int)
	SamplerParameterf(sampler Sampler, pname Enum, param float64)
	GetSamplerParameter(sampler Sampler, pname Enum) any

	/* Sync Methods */

	FenceSync(condition, flags Enum) Sync
	ClientWaitSync(sync Sync, flags Enum, timeout uint)
	WaitSync(sync Sync, flags Enum, timeout int)
	GetSyncParameter(sync Sync, pname Enum) any

	/* Transform Feedback Methods */

	BeginTransformFeedback(primitiveMode Enum)
	EndTransformFeedback()
	TransformFeedbackVaryings(program Program, varyings []string, bufferMode Enum)
	GetTransformFeedbackVarying(program Program, index uint) ActiveInfo
	PauseTransformFeedback()
	ResumeTransformFeedback()

	/* Transform Feedback / Uniform Buffer Object Methods */

	GetIndexedParameter(target Enum, index uint) any
	GetUniformIndices(program Program, uniformNames []string) []uint
	GetActiveUniforms(program Program, uniformIndices []uint, pname Enum) any
	GetUniformBlockIndex(program Program, uniformBlockName string) uint
	GetActiveUniformBlockParameter(program Program, uniformBlockIndex uint, pname Enum) any
	GetActiveUniformBlockName(program Program, uniformBlockIndex uint) string
	UniformBlockBinding(program Program, uniformBlockIndex, uniformBlockBinding uint)

	/* Viewport Methods */

	Viewport(x, y, width, height int)
	Scissor(x, y, width, height int)

	/* State Methods */

	BlendColor(red, green, blue, alpha float64)
	BlendEquation(mode Enum)
	BlendEquationSeparate(modeRGB, modeAlpha Enum)
	BlendFunc(sfactor, dfactor Enum)
	BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha Enum)
	Clear(mask Enum)
	ClearColor(red, green, blue, alpha float64)
	ClearDepth(depth float64)
	ClearStencil(s int)
	ColorMask(red, green, blue, alpha bool)
	CullFace(mode Enum)
	FrontFace(mode Enum)
	DepthFunc(fun Enum)
	DepthMask(flag bool)
	DepthRange(zNear, zFar float64)
	Disable(cap Enum)
	Enable(cap Enum)
	Hint(target, mode Enum)
	LineWidth(width float64)
	PixelStorei(pname Enum, param int)
	PolygonOffset(factor, units float64)
	StencilFunc(fun Enum, ref int, mask uint)
	StencilFuncSeparate(face, fun Enum, ref int, mask uint)
	StencilMask(mask uint)
	StencilMaskSeparate(face Enum, mask uint)
	StencilOp(fail, zfail, zpass Enum)
	StencilOpSeparate(face, fail, zfail, zpass Enum)
	SampleCoverage(value float64, invert bool)

	GetParameter(pname Enum) any

	/* Pipeline Control Methods */

	Finish()
	Flush()

	/* Draw Methods */

	DrawArrays(mode Enum, first, count int)
	DrawElements(mode Enum, count int, typ Enum, offset int)
	DrawArraysInstanced(mode Enum, first, count, instanceCount int)
	DrawElementsInstanced(mode Enum, count int, typ Enum, offset, instanceCount int)
	DrawRangeElements(mode Enum, start, end uint, count int, typ Enum, offset int)

	/* Error Methods */

	GetError() Enum
}
