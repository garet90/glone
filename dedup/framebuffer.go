package dedup

import "gfx.cafe/ghalliday1/glone"

type framebufferAttachment struct {
	alphaSize          int
	blueSize           int
	colorEncoding      glone.Enum
	componentType      glone.Enum
	depthSize          int
	greenSize          int
	objectName         any // glone.Renderbuffer or glone.Texture
	objectType         glone.Enum
	redSize            int
	stencilSize        int
	textureCubeMapFace int
	textureLayer       int
	textureLevel       int
}

type Framebuffer struct {
	glone.Framebuffer
}
