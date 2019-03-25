package view

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Canvas struct {
	parent     *Object
	projection mgl32.Mat4
	view       mgl32.Mat4
}

func NewCanvas(w *Window) *Canvas {
	return &Canvas{
		projection: mgl32.Ortho(float32(-w.Width/2), float32(w.Width/2), -float32(w.Height/2), float32(w.Height/2), 0, 100),
		view:       mgl32.LookAtV(mgl32.Vec3{0, 0, 50}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0}),
	}
}

func (c *Canvas) Init() {
	hW := float32(c.parent.scene.Window.Width / 2)
	hH := float32(c.parent.scene.Window.Height / 2)
	c.projection = mgl32.Ortho(-hW, hW, -hH, hH, 0, 100)
	c.view = mgl32.LookAtV(mgl32.Vec3{hW, -hH, 50}, mgl32.Vec3{hW, -hH, 0}, mgl32.Vec3{0, 1, 0})
}

func (c *Canvas) Update() {
	gl.UseProgram(spriteShader.program)
	projectionUniform := spriteShader.uniforms["projection"]
	viewUniform := spriteShader.uniforms["view"]

	gl.UniformMatrix4fv(projectionUniform, 1, false, &c.projection[0])
	gl.UniformMatrix4fv(viewUniform, 1, false, &c.view[0])
	for _, obj := range c.parent.Children {
		for _, comp := range obj.components {
			comp.Update()
		}
	}
}

func (c *Canvas) SetParent(obj *Object) {
	c.parent = obj
}
