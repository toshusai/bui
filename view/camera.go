package view

import "github.com/go-gl/mathgl/mgl32"

// Camera !
type Camera struct {
	projection mgl32.Mat4
	view       mgl32.Mat4
	position   mgl32.Vec3
}

// NewCamera create a new camera
func NewCamera() *Camera {
	return &Camera{
		projection: mgl32.Ortho(-400, 400, -300, 300, 100, -100),
		view:       mgl32.LookAtV(mgl32.Vec3{0, 0, -1}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0}),
		position:   mgl32.Vec3{0, 0, -1},
	}
}
