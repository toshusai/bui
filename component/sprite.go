package component

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/toshusai/bui/view"
)

// Sprite is 2D image
type Sprite struct {
	Texture       *view.Texture
	vertices      []float32
	vao           uint32
	vbo           uint32
	parent        *view.Object
	Shader        *view.Shader
	rectTransform *RectTransform
}

// NewSprite create a new sprite
func NewSprite(tex *view.Texture) *Sprite {
	sp := &Sprite{}
	sp.Shader = view.GetSpriteShader()
	sp.Texture = tex
	sp.vertices = []float32{
		0, 0, 0.0, 0, 0.0,
		0, float32(-sp.Texture.GetHeight()), 0.0, 0.0, 1.0,
		float32(sp.Texture.GetWidth()), 0, 0.0, 1.0, 0,

		0, float32(-sp.Texture.GetHeight()), 0.0, 0.0, 1.0,
		float32(sp.Texture.GetWidth()), float32(-sp.Texture.GetHeight()), 0.0, 1.0, 1.0,
		float32(sp.Texture.GetWidth()), 0, 0.0, 1.0, 0,
	}

	gl.GenVertexArrays(1, &sp.vao)
	gl.BindVertexArray(sp.vao)

	gl.GenBuffers(1, &sp.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, sp.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(sp.vertices)*4, gl.Ptr(sp.vertices), gl.DYNAMIC_DRAW)

	vertAttrib := uint32(sp.Shader.Uniforms["vert"])
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(sp.Shader.Uniforms["vertTexCoord"])
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	return sp
}

func (sp *Sprite) SetParent(obj *view.Object) {
	sp.parent = obj
}

func (sp *Sprite) Update() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.ONE, gl.ONE_MINUS_SRC_ALPHA)
	v := view.GetSpriteShader()
	gl.UseProgram(v.GetProgram())

	sp.vertices[6] = -sp.rectTransform.Height
	sp.vertices[10] = sp.rectTransform.Width
	sp.vertices[16] = -sp.rectTransform.Height
	sp.vertices[20] = sp.rectTransform.Width
	sp.vertices[21] = -sp.rectTransform.Height
	sp.vertices[25] = sp.rectTransform.Width

	gl.BindBuffer(gl.ARRAY_BUFFER, sp.vbo)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(sp.vertices)*4, gl.Ptr(sp.vertices))
	translate := mgl32.Translate3D(
		sp.parent.Position.X(),
		sp.parent.Position.Y(),
		sp.parent.Position.Z())

	scale := mgl32.Scale3D(
		sp.parent.Scale.X(),
		sp.parent.Scale.Y(),
		sp.parent.Scale.Z())

	model := translate.Mul4(scale)
	model = sp.parent.Rotation.Mul4(model)

	gl.UniformMatrix4fv(sp.Shader.Uniforms["model"], 1, false, &model[0])

	gl.BindVertexArray(sp.vao)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, sp.Texture.GetTexture())
	gl.DrawArrays(gl.TRIANGLES, 0, 6)

}

func (sp *Sprite) Init() {
	sp.rectTransform = sp.parent.GetComponent(&RectTransform{}).(*RectTransform)

}
