package view

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Object struct {
	Position   mgl32.Vec3
	Rotation   mgl32.Mat4
	Scale      mgl32.Vec3
	components []component
	scene      *Scene
	Parent     *Object
	Children   []*Object
}

func NewObject() *Object {
	return &Object{
		Position: mgl32.Vec3{},
		Rotation: mgl32.Ident4(),
		Scale:    mgl32.Vec3{1, 1, 1},
	}
}

func (obj *Object) AddComponent(comp component) {
	comp.SetParent(obj)
	obj.components = append(obj.components, comp)
}
