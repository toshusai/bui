package view

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Scene !
type Scene struct {
	objects []*Object
	Window  *Window
}

var currentScene *Scene

func GetScene() *Scene {
	return currentScene
}

// NewScene create a new scene
func NewScene() *Scene {
	var err error
	if err != nil {
		panic(err)
	}

	scene := &Scene{}
	currentScene = scene
	return scene
}

// Add sprite to scene
func (scene *Scene) Add(obj *Object) {
	scene.objects = append(scene.objects, obj)
}

func (scene *Scene) Start() {
	for _, obj := range scene.objects {
		obj.Init()
	}
}

// Draw all objects
func (scene *Scene) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	for _, obj := range scene.objects {
		obj.Update()
	}
}
