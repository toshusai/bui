package view

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Scene !
type Scene struct {
	Camera  *Camera
	objects []*Object
	Window  *Window
}

// NewScene create a new scene
func NewScene() *Scene {
	var err error
	program, err = newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	scene := &Scene{}

	return scene
}

// Add sprite to scene
func (scene *Scene) Add(obj *Object) {
	obj.scene = scene
	scene.objects = append(scene.objects, obj)
}

// Draw all objects
func (scene *Scene) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	for _, obj := range scene.objects {
		for _, comp := range obj.components {
			comp.Update()
		}
	}
}
