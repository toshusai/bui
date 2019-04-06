package component

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/toshusai/bui/view"
)

type Square struct {
	vertices      []float32
	vao           uint32
	vbo           uint32
	parent        *view.Object
	Shader        *view.Shader
	rectTransform *RectTransform
}

func NewSquare() *Square {
	sq := &Square{}
	sq.vertices = []float32{
		0, 0, 0,
		0, -1, 0,
		1, 0, 0,

		0, -1, 0,
		1, -1, 0,
		1, 0, 0,
	}

	gl.GenVertexArrays(1, &sq.vao)
	gl.BindVertexArray(sq.vao)
	gl.GenBuffers(1, &sq.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, sq.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(sq.vertices)*4, gl.Ptr(sq.vertices), gl.DYNAMIC_DRAW)

	sq.Shader = view.GetSimpleShader()
	vertAttrib := uint32(sq.Shader.Uniforms["vert"])
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	return sq
}

func (sq *Square) SetParent(obj *view.Object) {
	sq.parent = obj
}

func (sq *Square) Init() {
	sq.rectTransform = sq.parent.GetComponent(&RectTransform{}).(*RectTransform)
}

func (sq *Square) Update() {
	v := view.GetSimpleShader()
	gl.UseProgram(v.GetProgram())

	sq.vertices[4] = -sq.rectTransform.Height
	sq.vertices[6] = sq.rectTransform.Width
	sq.vertices[10] = -sq.rectTransform.Height
	sq.vertices[12] = sq.rectTransform.Width
	sq.vertices[13] = -sq.rectTransform.Height
	sq.vertices[15] = sq.rectTransform.Width

	gl.BindBuffer(gl.ARRAY_BUFFER, sq.vbo)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(sq.vertices)*4, gl.Ptr(sq.vertices))

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
