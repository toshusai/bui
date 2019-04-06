package main

import (
	"fmt"
	"log"

	"github.com/go-gl/mathgl/mgl32"

	_ "image/png"

	"github.com/toshusai/bui"
	"github.com/toshusai/bui/component"
	"github.com/toshusai/bui/view"
)

func main() {
	count := 0

	w := view.NewWindow(800, 600, "Test")
	scene := view.NewScene()
	w.AddScene(scene)
	view.InitShader()

	// Create Object (Canvas)
	can := component.NewCanvas(w)
	canObj := view.NewObject()
	canObj.AddComponent(can)

	// Create Texture
	tex, err := view.NewTexture("test_image_32px.png")
	if err != nil {
		log.Fatalln(err)
	}

	// Create Font
	f, e := view.LoadFont("C:/Windows/Fonts/meiryo.ttc", 32, 800, 600)
	if e != nil {
		panic(e)
	}

	// Create Object (Text)
	textObj := view.NewObject()
	textObj.Position = mgl32.Vec3{200, 200, 0}
	text := component.NewText()
	textObj.AddComponent(text)
	text.Size = 1
	text.Font = f
	text.Color = bui.Color{1, 0, 0, 1}
	canObj.AddChild(textObj)

	// Create Object (Sprite, Button)
	sp := component.NewSprite(tex)
	spObj := view.NewObject()
	spObj.Position = mgl32.Vec3{32, -32, 0}
	spObj.AddComponent(sp)

	btn := component.NewButton()
	btn.OnClick = func() {
		count++
		text.Text = fmt.Sprintf("+%d", count)
	}
	spObj.AddComponent(btn)
	canObj.AddChild(spObj)

	// Create Object (Square)
	sqObj := view.NewObject()
	sqObj.Position = mgl32.Vec3{100, -100, 0}
	sq := component.NewSquare()
	sqObj.AddComponent(sq)
	canObj.AddChild(sqObj)

	scene.Add(canObj)

	scene.Start()
	w.Update = func() {
		scene.Draw()
	}
	w.Run()
}
