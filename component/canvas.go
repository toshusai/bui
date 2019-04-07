package component

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/toshusai/bui/view"
)

type Canvas struct {
	parent        *view.Object
	projection    mgl32.Mat4
	view          mgl32.Mat4
	rectTransform *RectTransform
}

func NewCanvas(w *view.Window) *Canvas {
	return &Canvas{}
}

func (c *Canvas) Init() {
	c.rectTransform = c.parent.GetComponent(&RectTransform{}).(*RectTransform)
	hW := float32(c.rectTransform.Width / 2)
	hH := float32(c.rectTransform.Height / 2)
	c.projection = mgl32.Ortho(-hW, hW, -hH, hH, 0, 100)
	c.view = mgl32.LookAtV(mgl32.Vec3{hW, -hH, 50}, mgl32.Vec3{hW, -hH, 0}, mgl32.Vec3{0, 1, 0})
}

func (c *Canvas) Update() {
	spriteShader := view.GetSpriteShader()
	simpleShader := view.GetSimpleShader()
	gl.UseProgram(spriteShader.GetProgram())
	gl.UniformMatrix4fv(spriteShader.Uniforms["projection"], 1, false, &c.projection[0])
	gl.UniformMatrix4fv(spriteShader.Uniforms["view"], 1, false, &c.view[0])
	gl.UseProgram(simpleShader.GetProgram())
	gl.UniformMatrix4fv(simpleShader.Uniforms["projection"], 1, false, &c.projection[0])
	gl.UniformMatrix4fv(simpleShader.Uniforms["view"], 1, false, &c.view[0])
}

func (c *Canvas) SetParent(obj *view.Object) {
	c.parent = obj
}
