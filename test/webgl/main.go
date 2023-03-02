//go:build js && wasm

package main

import (
	_ "embed"
	"encoding/binary"
	"fmt"
	"github.com/garet90/glone"
	"github.com/garet90/glone/dedup"
	"github.com/garet90/glone/webgl"
	"math"
	"strings"
	"syscall/js"
)

//go:embed shader.vert
var vertexShaderSrc string

//go:embed shader.frag
var fragmentShaderSrc string

func main() {
	window := js.Global()
	document := window.Get("document")
	body := document.Get("body")
	canvas := document.Call("createElement", "canvas")
	body.Call("appendChild", canvas)
	renderingContext := canvas.Call("getContext", "webgl2")
	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()

	var gl glone.RenderingContext
	gl = webgl.NewRenderingContext(renderingContext)
	gl = dedup.NewRenderingContext(gl)

	// create shaders
	vertex := gl.CreateShader(glone.VERTEX_SHADER)
	fragment := gl.CreateShader(glone.FRAGMENT_SHADER)
	gl.ShaderSource(vertex, vertexShaderSrc)
	gl.CompileShader(vertex)
	gl.ShaderSource(fragment, fragmentShaderSrc)
	gl.CompileShader(fragment)

	// create program
	program := gl.CreateProgram()
	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)
	gl.LinkProgram(program)
	gl.ValidateProgram(program)
	if gl.GetProgramParameter(program, glone.LINK_STATUS).(bool) != true {
		var msg strings.Builder
		programLog := gl.GetProgramInfoLog(program)
		msg.WriteString("Program Log: ")
		msg.WriteString(programLog)
		msg.WriteRune('\n')
		if gl.GetShaderParameter(vertex, glone.COMPILE_STATUS).(bool) != true {
			vertexLog := gl.GetShaderInfoLog(vertex)
			msg.WriteString("Vertex Log: ")
			msg.WriteString(vertexLog)
			msg.WriteRune('\n')
		}
		if gl.GetShaderParameter(fragment, glone.COMPILE_STATUS).(bool) != true {
			fragmentLog := gl.GetShaderInfoLog(fragment)
			msg.WriteString("Fragment Log: ")
			msg.WriteString(fragmentLog)
			msg.WriteRune('\n')
		}
		panic(msg.String())
	}
	gl.UseProgram(program)

	vertexPositionLocation := gl.GetAttribLocation(program, "vertex_position")
	if vertexPositionLocation < 0 {
		panic("couldn't find vertex_position")
	}
	vertexColorLocation := gl.GetAttribLocation(program, "vertex_color")
	if vertexColorLocation < 0 {
		panic("couldn't find vertex_color")
	}

	var b []byte
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(-0.5))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(-0.5))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(1.0))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0.0))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0.0))

	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0.5))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0.0))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(1.0))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0.0))

	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0.5))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(-0.5))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0.0))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(0.0))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(1.0))

	// create buffer
	buffer := gl.CreateBuffer()
	gl.BindBuffer(glone.ARRAY_BUFFER, buffer)
	gl.BufferDataPix(glone.ARRAY_BUFFER, b, glone.STREAM_DRAW)

	// create vao
	vao := gl.CreateVertexArray()
	gl.BindVertexArray(vao)
	gl.VertexAttribPointer(uint32(vertexPositionLocation), 2, glone.FLOAT, false, 20, 0)
	gl.VertexAttribPointer(uint32(vertexColorLocation), 3, glone.FLOAT, false, 20, 8)
	gl.EnableVertexAttribArray(uint32(vertexPositionLocation))
	gl.EnableVertexAttribArray(uint32(vertexColorLocation))

	// clear canvas
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(glone.COLOR_BUFFER_BIT | glone.DEPTH_BUFFER_BIT)

	gl.Viewport(0, 0, int32(width), int32(height))

	// draw!
	gl.DrawArrays(glone.TRIANGLES, 0, 3)

	err := gl.GetError()
	if err != glone.NO_ERROR {
		panic(fmt.Sprintf("gl error: %d", err))
	}
}
