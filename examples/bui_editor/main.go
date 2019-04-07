package main

import (
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/toshusai/bui"
	"github.com/toshusai/bui/component"
	"github.com/toshusai/bui/view"
)

var font *view.Font
var scene *view.Scene
var window *view.Window

func initFont() {
	var e error
	font, e = view.LoadFont("C:/Windows/Fonts/meiryo.ttc", 32, 800, 600)
	if e != nil {
		panic(e)
	}
}

func Start() {
	scene.Start()
	window.Update = func() {
		scene.Draw()
	}
	window.Run()
}

func main() {
	window = view.NewWindow(800, 600, "Test")
	scene = view.NewScene()
	window.AddScene(scene)
	view.InitShader()
	initFont()

	// Create Object (Canvas)
	canObj := view.NewObject()
	canObj.Name = "canvas"

	can := component.NewCanvas(window)

	canRT := &component.RectTransform{}
	canRT.Width = 800
	canRT.Height = 600

	canObj.AddComponent(can)
	canObj.AddComponent(canRT)

	// Create Object (Text)
	textObj := view.NewObject()
	textObj.Name = "text"
	textObj.Position = mgl32.Vec3{0, 0, 0}

	textSq := component.NewSquare()
	textObj.AddComponent(textSq)

	text := component.NewText()
	text.Size = 1
	text.Font = font
	text.Color = bui.Color{0.5, 0.7, 0.3, 1}
	textObj.AddComponent(text)

	textRt := component.NewRectTransform()
	textRt.Width = 100
	textRt.Height = 100
	textObj.AddComponent(textRt)

	canObj.AddChild(textObj)

	scene.Add(canObj)
	text.Text = "\n" + s(scene.Root)
	Start()
}

var depth = 0

func s(obj *view.Object) string {
	str := ""
	if len(obj.Children) == 0 {
		return ""
	}
	for _, obj := range obj.Children {
		for i := 0; i < depth; i++ {
			str += " "
		}
		str += obj.Name + "\n"
		for _, comp := range obj.Components {
			for i := 0; i < depth; i++ {
				str += " "
			}
			str += " "
			str += reflect.TypeOf(comp).String() + "\n"
		}
		depth++
		str += s(obj)
	}
	return str
}
