package main

import (
	"log"

	"github.com/toshusai/bui"

	"github.com/go-gl/mathgl/mgl32"

	_ "image/png"

	"github.com/toshusai/bui/component"
	"github.com/toshusai/bui/view"
)

func main() {

	w := view.NewWindow(800, 600, "Test")
	scene := view.NewScene()
	w.AddScene(scene)
	view.InitShader()

	// Create texture
	tex, err := view.NewTexture("test_image_32px.png")
	if err != nil {
		log.Fatalln(err)
	}

	f, e := view.LoadFont("C:/Windows/Fonts/meiryo.ttc", 50, 800, 600)
	if e != nil {
		panic(e)
	}

	// Create Object (Sprite, Button)
	sp := component.NewSprite(tex)
	obj := view.NewObject()
	obj.Position = mgl32.Vec3{32, -32, 0}
	obj.AddComponent(sp)

	btn := component.NewButton()
	btn.OnClick = func() {
		obj.Position = obj.Position.Add(mgl32.Vec3{1, 1, 0})
	}
	obj.AddComponent(btn)

	// Create Object (Canvas)
	can := component.NewCanvas(w)
	canObj := view.NewObject()
	canObj.AddChild(obj)
	canObj.AddComponent(can)
	btn.Init()

	// Create Object (Square)
	sqObj := view.NewObject()
	sqObj.Position = mgl32.Vec3{100, -100, 0}
	sq := component.NewSquare()
	sqObj.AddComponent(sq)
	canObj.AddChild(sqObj)

	// Create Object (Text)
	textObj := view.NewObject()
	textObj.Position = mgl32.Vec3{200, 200, 0}
	text := component.NewText()
	textObj.AddComponent(text)
	text.Size = 1
	text.Text = "Hey"
	text.Font = f
	text.Color = bui.Color{1, 0, 0, 1}
	canObj.AddChild(textObj)

	// cam := component.NewCamera()
	// camObj := view.NewObject()
	// camObj.Position = mgl32.Vec3{0, 0, -1}
	// camObj.AddComponent(cam)

	scene.Add(canObj)
	// scene.Add(camObj)

	scene.Start()
	w.Update = func() {
		scene.Draw()
	}
	w.Run()
}
