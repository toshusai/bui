package view

import (
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
)

type Object struct {
	Position   mgl32.Vec3
	Rotation   mgl32.Mat4
	Scale      mgl32.Vec3
	components []component
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

func (obj *Object) AddChild(chld *Object) {
	obj.Children = append(obj.Children, chld)
}

func (obj *Object) Init() {
	for _, child := range obj.Children {
		child.Init()
	}
	for _, comp := range obj.components {
		comp.Init()
	}
}

func (obj *Object) Update() {
	for _, child := range obj.Children {
		child.Update()
	}
	for _, comp := range obj.components {
		comp.Update()
	}
}

func (obj *Object) GetComponent(value interface{}) component {
	for _, comp := range obj.components {
		if reflect.TypeOf(value) == reflect.TypeOf(comp) {
			return comp
		}
	}
	return nil
}
