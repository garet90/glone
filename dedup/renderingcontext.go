package dedup

import (
	"gfx.cafe/ghalliday1/glone"
	"math"
)

type RenderingContext struct {
	glone.RenderingContext

	// buffers

	activeTexture                             glone.Enum
	aliasedLineWidthRange                     [2]float32
	aliasedPointSizeRange                     [2]float32
	alphaBits                                 int
	arrayBufferBinding                        glone.Buffer
	blend                                     bool
	blendColor                                [4]float32
	blendDstAlpha                             glone.Enum
	blendDstRGB                               glone.Enum
	blendEquationAlpha                        glone.Enum
	blendEquationRGB                          glone.Enum
	blendSrcAlpha                             glone.Enum
	blendSrcRGB                               glone.Enum
	blueBits                                  int
	colorClearValue                           [4]float32
	colorWritemask                            [4]bool
	compressedTextureFormats                  []uint32
	cullFace                                  bool
	cullFaceMode                              glone.Enum
	currentProgram                            glone.Program
	depthBits                                 int
	depthClearValue                           float32
	depthFunc                                 glone.Enum
	depthRange                                [2]float32
	depthTest                                 bool
	depthWritemask                            bool
	dither                                    bool
	elementArrayBufferBinding                 glone.Buffer
	framebufferBinding                        glone.Framebuffer
	frontFace                                 glone.Enum
	generateMipmapHint                        glone.Enum
	greenBits                                 int
	implementationColorReadFormat             glone.Enum
	implementationColorReadType               glone.Enum
	lineWidth                                 float32
	maxCombinedTextureImageUnits              int
	maxCubeMapTextureSize                     int
	maxFragmentUniformVectors                 int
	mexRenderbufferSize                       int
	maxTextureImageUnits                      int
	maxTextureSize                            int
	maxVaryingVectors                         int
	maxVertexAttribs                          int
	maxVertexTextureImageUnits                int
	maxVertexUniformVectors                   int
	maxViewportDims                           [2]int32
	packAlignment                             int
	polygonOffsetFactor                       float32
	polygonOffsetFill                         bool
	polygonOffsetUnits                        float32
	redBits                                   int
	renderbufferBinding                       glone.Renderbuffer
	renderer                                  string
	sampleAlphaToCoverage                     bool
	sampleBuffers                             int
	sampleCoverage                            bool
	sampleCoverageInvert                      bool
	sampleCoverageValue                       bool
	samples                                   int
	scissorBox                                [4]int32
	scissorTest                               bool
	shadingLanguageVersion                    string
	stencilBackFail                           glone.Enum
	stencilBackFunc                           glone.Enum
	stencilBackPassDepthFail                  glone.Enum
	stencilBackPassDepthPass                  glone.Enum
	stencilBackRef                            int
	stencilBackValueMask                      uint
	stencilBackWriteMask                      uint
	stencilBits                               int
	stencilClearValue                         int
	stencilFail                               glone.Enum
	stencilFunc                               glone.Enum
	stencilPassDepthFail                      glone.Enum
	stencilPassDepthPass                      glone.Enum
	stencilRef                                int
	stencilTest                               bool
	stencilValueMask                          uint
	stencilWritemask                          uint
	subpixelBits                              int
	textureBinding2D                          glone.Texture
	textureBindingCubeMap                     glone.Texture
	unpackAlignment                           int
	unpackColorspaceConversionWebgl           glone.Enum
	unpackFlipYWebgl                          bool
	unpackPremultiplyAlphaWebgl               bool
	vendor                                    string
	version                                   string
	viewport                                  [4]int32
	copyReadBufferBinding                     glone.Buffer
	copyWriteBufferBinding                    glone.Buffer
	drawBuffers                               []glone.Enum
	drawFramebufferBinding                    glone.Framebuffer
	fragmentShaderDerivativeHint              glone.Enum
	max3DTextureSize                          int
	maxArrayTextureLayers                     int
	maxClientWaitTimeoutWebgl                 int64
	maxColorAttachments                       int
	maxCombinedFragmentUniformComponents      int64
	maxCombinedUniformBlocks                  int
	maxCombinedVertexUniformComponents        int64
	maxDrawBuffers                            int
	maxElementIndex                           int64
	maxElementsIndices                        int
	maxElementsVertices                       int
	maxFragmentInputComponents                int
	maxFragmentUniformBlocks                  int
	maxFragmentUniformComponents              int
	maxProgramTexelOffset                     int
	maxSamples                                int
	maxServerWaitTimeout                      int64
	maxTextureLodBias                         float32
	maxTransformFeedbackInterleavedComponents int
	maxTransformFeedbackSeparateAttribs       int
	maxTransformFeedbackSeparateComponents    int
	maxUniformBlockSize                       int64
	uniformBufferBindings                     int
	maxVaryingComponents                      int
	maxVertexOutputComponents                 int
	maxVertexUniformBlocks                    int
	maxVertexUniformComponents                int
	minProgramTexelOffset                     int
	packRowLength                             int
	packSkipPixels                            int
	packSkipRows                              int
	pixelPackBufferBinding                    glone.Buffer
	pixelUnpackBufferBinding                  glone.Buffer
	rasterizerDiscard                         bool
	readBuffer                                glone.Enum
	readFramebufferBinding                    glone.Framebuffer
	samplerBinding                            glone.Sampler
	textureBinding2DArray                     glone.Texture
	textureBinding3D                          glone.Texture
	transformFeedbackActive                   bool
	transformFeedbackBinding                  glone.TransformFeedback
	transformFeedbackBufferBinding            glone.Buffer
	transformFeedbackPaused                   bool
	uniformBufferBinding                      glone.Buffer
	uniformBufferOffsetAlignment              int
	unpackImageHeight                         int
	unpackRowLength                           int
	unpackSkipImages                          int
	unpackSkipPixels                          int
	unpackSkipRows                            int
	vertexArrayBinding                        glone.VertexArray
	// TODO getIndexedParameter
}

func MakeRenderingContext(context glone.RenderingContext) RenderingContext {
	return RenderingContext{
		RenderingContext: context,

		activeTexture:             glone.TEXTURE0,
		arrayBufferBinding:        nil,
		blend:                     false,
		blendDstAlpha:             glone.ZERO,
		blendDstRGB:               glone.ZERO,
		blendSrcAlpha:             glone.ONE,
		blendSrcRGB:               glone.ONE,
		colorClearValue:           [4]float32{0, 0, 0, 0},
		colorWritemask:            [4]bool{true, true, true, true},
		cullFace:                  false,
		cullFaceMode:              glone.BACK,
		depthClearValue:           1.0,
		depthFunc:                 glone.LESS,
		depthRange:                [2]float32{0, 1},
		depthTest:                 false,
		depthWritemask:            true,
		dither:                    true,
		elementArrayBufferBinding: nil,
		framebufferBinding:        nil,
		frontFace:                 glone.CCW,
		generateMipmapHint:        glone.DONT_CARE,
		lineWidth:                 1,
		packAlignment:             4,
		polygonOffsetFactor:       0,
		polygonOffsetFill:         false,
		polygonOffsetUnits:        0,
		renderbufferBinding:       nil,
		sampleAlphaToCoverage:     false,
		sampleCoverage:            false,
		scissorTest:               false,
		stencilBackFail:           glone.KEEP,
		stencilBackFunc:           glone.ALWAYS,
		stencilBackPassDepthFail:  glone.KEEP,
		stencilBackPassDepthPass:  glone.KEEP,
		stencilBackRef:            0,
		stencilBackValueMask:      math.MaxUint,
		stencilClearValue:         0,
		stencilFail:               glone.KEEP,
		stencilFunc:               glone.ALWAYS,
		stencilPassDepthFail:      glone.KEEP,
		stencilPassDepthPass:      glone.KEEP,
		stencilRef:                0,
		stencilTest:               false,
		stencilValueMask:          math.MaxUint,
		stencilWritemask:          math.MaxUint,
		textureBinding2D:          nil,
		textureBindingCubeMap:     nil,
	}
}

func NewRenderingContext(context glone.RenderingContext) *RenderingContext {
	rc := MakeRenderingContext(context)
	return &rc
}

func (R *RenderingContext) BindBuffer(target glone.Enum, buffer glone.Buffer) {
	switch target {
	case glone.ARRAY_BUFFER:
		if buffer == R.arrayBufferBinding {
			return
		}
		R.arrayBufferBinding = buffer
	case glone.ELEMENT_ARRAY_BUFFER:
		if buffer == R.elementArrayBufferBinding {
			return
		}
		R.elementArrayBufferBinding = buffer
	case glone.COPY_READ_BUFFER:
		if buffer == R.copyReadBufferBinding {
			return
		}
		R.copyReadBufferBinding = buffer
	case glone.COPY_WRITE_BUFFER:
		if buffer == R.copyWriteBufferBinding {
			return
		}
		R.copyWriteBufferBinding = buffer
	case glone.TRANSFORM_FEEDBACK_BUFFER:
		if buffer == R.transformFeedbackBufferBinding {
			return
		}
		R.transformFeedbackBufferBinding = buffer
	case glone.UNIFORM_BUFFER:
		if buffer == R.uniformBufferBinding {
			return
		}
		R.uniformBufferBinding = buffer
	case glone.PIXEL_PACK_BUFFER:
		if buffer == R.pixelPackBufferBinding {
			return
		}
		R.pixelPackBufferBinding = buffer
	case glone.PIXEL_UNPACK_BUFFER:
		if buffer == R.pixelUnpackBufferBinding {
			return
		}
		R.pixelUnpackBufferBinding = buffer
	default:
		return
	}
	R.RenderingContext.BindBuffer(target, buffer)
}

func (R *RenderingContext) ActiveTexture(texture glone.Enum) {
	if texture == R.activeTexture {
		return
	}
	R.activeTexture = texture
	R.RenderingContext.ActiveTexture(texture)
}
