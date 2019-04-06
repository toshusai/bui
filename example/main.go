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

type Drag struct {
}

func NewDragObject(w *view.Window, P *view.Object) *view.Object {
	obj := view.NewObject()
	obj.Position = mgl32.Vec3{300, -100, 0}

	sq := component.NewSquare()
	obj.AddComponent(sq)

	rt := &component.RectTransform{}
	rt.Width = 50
	rt.Height = 50
	obj.AddComponent(rt)

	btn := component.NewButton()
	trt := P.GetComponent(&component.RectTransform{}).(*component.RectTransform)
	btn.OnDrag = func() {
		x, y := w.GetCursorPos()
		obj.Position = mgl32.Vec3{float32(x) - rt.Width/2, -float32(y) + rt.Height/2, 0}
		trt.Width = float32(x) - P.Position.X()
		trt.Height = float32(y) + P.Position.Y()
	}
	obj.AddComponent(btn)

	return obj
}

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

	// Create Object (Square)
	sqObj := view.NewObject()
	sqObj.Position = mgl32.Vec3{100, -100, 0}
	sq := component.NewSquare()
	rt := &component.RectTransform{}
	rt.Width = 50
	rt.Height = 50
	sqObj.AddComponent(rt)
	sqObj.AddComponent(sq)
	canObj.AddChild(sqObj)

	// Create Object (Sprite, Button)
	sp := component.NewSprite(tex)
	spObj := view.NewObject()
	spObj.Position = mgl32.Vec3{32, -32, 0}
	spObj.AddComponent(sp)

	rt2 := &component.RectTransform{}
	rt2.Width = 50
	rt2.Height = 50
	spObj.AddComponent(rt2)

	btn := component.NewButton()
	btn.OnPointerDwon = func() {
		count++
		text.Text = fmt.Sprintf("+%d", count)
		rt.Width += 2
		rt2.Width += 2
	}
	btn.OnDrag = func() {
		x, y := w.GetCursorPos()
		spObj.Position = mgl32.Vec3{float32(x) - rt2.Width/2, -float32(y) + rt2.Height/2, 0}
	}
	btn.OnPointerDwon = func() {
		fmt.Println("Hey")
	}
	spObj.AddComponent(btn)
	canObj.AddChild(spObj)

	d := NewDragObject(w, spObj)
	canObj.AddChild(d)

	scene.Add(canObj)

	scene.Start()
	w.Update = func() {
		scene.Draw()
	}
	w.Run()
}
