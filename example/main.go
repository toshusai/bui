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
	view.InitShader()

	tex, err := view.NewTexture("test_image_32px.png")
	if err != nil {
		log.Fatalln(err)
	}

	sp := view.NewSprite(tex)
	obj := view.NewObject()
	obj.Position = mgl32.Vec3{0, 0, 0}
	obj.AddComponent(sp)

	can := view.NewCanvas(w)
	canObj := view.NewObject()
	canObj.Children = []*view.Object{obj}
	canObj.AddComponent(can)
	scene.Add(canObj)

	cam := view.NewCamera()
	camObj := view.NewObject()
	camObj.Position = mgl32.Vec3{0, 0, -1}
	camObj.AddComponent(cam)
	scene.Add(camObj)
	scene.Camera = cam

	w.Update = func() {
		scene.Draw()
		obj.Position = obj.Position.Add(mgl32.Vec3{1, 0, 0})
	}
	w.Run()
}
