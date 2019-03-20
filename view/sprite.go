package view

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Sprite is 2D image
type Sprite struct {
	Texture  *Texture
	vertices []float32
	vao      uint32
	program  uint32
	Scale    mgl32.Vec3
	parent   *Object
}

var program uint32

// NewSprite create a new sprite
func NewSprite(tex *Texture) *Sprite {
	sp := &Sprite{}
	sp.Texture = tex
	sp.vertices = []float32{
		0, 0, 0.0, 0, 0.0,
		0, float32(-sp.Texture.height), 0.0, 0.0, 1.0,
		float32(sp.Texture.width), 0, 0.0, 1.0, 0,

		0, float32(-sp.Texture.height), 0.0, 0.0, 1.0,
		float32(sp.Texture.width), float32(-sp.Texture.height), 0.0, 1.0, 1.0,
		float32(sp.Texture.width), 0, 0.0, 1.0, 0,
	}

	gl.GenVertexArrays(1, &sp.vao)
	gl.BindVertexArray(sp.vao)
	sp.program = program
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(sp.vertices)*4, gl.Ptr(sp.vertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(sp.program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(sp.program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	return sp
}

func (sp *Sprite) SetParent(obj *Object) {
	sp.parent = obj
}

func (sp *Sprite) Update() {
	gl.UseProgram(sp.program)
	projectionUniform := gl.GetUniformLocation(sp.program, gl.Str("projection\x00"))
	viewUniform := gl.GetUniformLocation(sp.program, gl.Str("view\x00"))

	gl.UniformMatrix4fv(projectionUniform, 1, false, &sp.parent.scene.Camera.projection[0])
	gl.UniformMatrix4fv(viewUniform, 1, false, &sp.parent.scene.Camera.view[0])

	translate := mgl32.Translate3D(
		sp.parent.Position.X(),
		sp.parent.Position.Y(),
		sp.parent.Position.Z())

	scale := mgl32.Scale3D(
		sp.parent.Scale.X(),
		sp.parent.Scale.Y(),
		sp.parent.Scale.Z())

	// sp.parent.Rotation = mgl32.LookAt(
	// 	0, 0, 0,
	// 	-sp.parent.scene.Camera.parent.Position.X(),
	// 	-sp.parent.scene.Camera.parent.Position.Y(),
	// 	-sp.parent.scene.Camera.parent.Position.Z(),
	// 	0, 1, 0)
	// sp.parent.Rotation = sp.parent.Rotation.Inv()

	model := translate.Mul4(scale)
	model = sp.parent.Rotation.Mul4(model)

	modelUniform := gl.GetUniformLocation(sp.program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	gl.BindVertexArray(sp.vao)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, sp.Texture.texture)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}
