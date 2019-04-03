package component

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/toshusai/bui/view"
)

type Square struct {
	vertices []float32
	vao      uint32
	parent   *view.Object
	Shader   *view.Shader
}

func NewSquare() *Square {
	sq := &Square{}
	sq.vertices = []float32{
		0, 0, 0,
		0, -50, 0,
		50, 0, 0,

		0, -50, 0,
		50, -50, 0,
		50, 0, 0,
	}

	gl.GenVertexArrays(1, &sq.vao)
	gl.BindVertexArray(sq.vao)
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(sq.vertices)*4, gl.Ptr(sq.vertices), gl.STATIC_DRAW)

	sq.Shader = view.GetSimpleShader()
	vertAttrib := uint32(sq.Shader.Uniforms["vert"])
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	return sq
}

func (sq *Square) SetParent(obj *view.Object) {
	sq.parent = obj
}

func (sq *Square) Update() {
	v := view.GetSimpleShader()
	gl.UseProgram(v.GetProgram())

	translate := mgl32.Translate3D(
		sq.parent.Position.X(),
		sq.parent.Position.Y(),
		sq.parent.Position.Z())

	scale := mgl32.Scale3D(
		sq.parent.Scale.X(),
		sq.parent.Scale.Y(),
		sq.parent.Scale.Z())

	model := translate.Mul4(scale)
	model = sq.parent.Rotation.Mul4(model)

	gl.UniformMatrix4fv(sq.Shader.Uniforms["model"], 1, false, &model[0])

	gl.BindVertexArray(sq.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func (sq *Square) Init() {

}
