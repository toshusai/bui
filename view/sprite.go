package view

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Sprite is 2D image
type Sprite struct {
	Texture  *Texture
	vertices []float32
	vao      uint32
	parent   *Object
	shader   *Shader
}

// NewSprite create a new sprite
func NewSprite(tex *Texture) *Sprite {
	sp := &Sprite{}
	sp.shader = spriteShader
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
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(sp.vertices)*4, gl.Ptr(sp.vertices), gl.STATIC_DRAW)

	vertAttrib := uint32(sp.shader.uniforms["vert"])
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(sp.shader.uniforms["vertTexCoord"])
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	return sp
}

func (sp *Sprite) SetParent(obj *Object) {
	sp.parent = obj
}

func (sp *Sprite) Update() {

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

	gl.UniformMatrix4fv(sp.shader.uniforms["model"], 1, false, &model[0])

	gl.BindVertexArray(sp.vao)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, sp.Texture.texture)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func (sp *Sprite) Init() {

}
