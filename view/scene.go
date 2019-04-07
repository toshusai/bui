package view

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Scene !
type Scene struct {
	Root   *Object
	Window *Window
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

	root := NewObject()
	root.Name = "root"
	scene := &Scene{
		Root: root,
	}
	currentScene = scene
	return scene
}

// Add sprite to scene
func (scene *Scene) Add(obj *Object) {
	scene.Root.AddChild(obj)
}

func (scene *Scene) Start() {
	for _, obj := range scene.Root.Children {
		obj.Init()
	}
}

// Draw all objects
func (scene *Scene) Draw() {
	// gl.ClearColor(1.0, 1.0, 0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	for _, obj := range scene.Root.Children {
		obj.Update()
	}
}
