package main

import (
	"log"

	_ "image/png"

	"github.com/toshusai/bui/view"
)

func main() {
	w := view.NewWindow(800, 600, "Test")

	scene := view.NewScene()
	scene.Camera = view.NewCamera()

	tex, err := view.NewTexture("test_image_32px.png")
	if err != nil {
		log.Fatalln(err)
	}

	sp := view.NewSprite(tex)

	obj := view.NewObject()

	obj.AddComponent(sp)

	scene.Add(&obj)

	w.Update = func() {
		scene.Draw()
	}
	w.Run()
}
