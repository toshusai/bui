package main

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"

	_ "image/png"

	"github.com/toshusai/bui/view"
)

func main() {
	w := view.NewWindow(800, 600, "Test")

	scene := view.NewScene()

	tex, err := view.NewTexture("test_image_32px.png")
	if err != nil {
		log.Fatalln(err)
	}

	sp := view.NewSprite(tex)
	obj := view.NewObject()
	obj.Position = mgl32.Vec3{50, 0, 0}
	obj.AddComponent(sp)
	scene.Add(obj)

	cam := view.NewCamera()
	camObj := view.NewObject()
	camObj.Position = mgl32.Vec3{0, 0, -1}
	camObj.AddComponent(cam)
	scene.Add(camObj)
	scene.Camera = cam

	w.Update = func() {
		scene.Draw()
	}
	w.Run()
}
