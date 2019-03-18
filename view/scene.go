package view

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Scene !
type Scene struct {
	Camera  *Camera
	sprites []*Sprite
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
func (scene *Scene) Add(sp *Sprite) {
	scene.sprites = append(scene.sprites, sp)
	sp.scene = scene
}

// Draw all sprites
func (scene *Scene) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	for _, sp := range scene.sprites {
		sp.draw()
	}
}
