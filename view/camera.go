package view

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Camera !
type Camera struct {
	projection mgl32.Mat4
	view       mgl32.Mat4
	parent     *Object
}

// NewCamera create a new camera
func NewCamera() *Camera {
	return &Camera{
		projection: mgl32.Perspective(mgl32.DegToRad(90.0), 800.0/600.0, 0.1, 100),
		view:       mgl32.LookAtV(mgl32.Vec3{0, 0, 50}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0}),
	}
}

func (cam *Camera) SetParent(obj *Object) {
	cam.parent = obj
}

func (cam *Camera) Update() {

}

func (cam *Camera) Init() {}
