package view

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type Button struct {
	OnClick func()
	parent  *Object
	sprite  *Sprite
}

func NewButton() *Button {
	return &Button{
		OnClick: func() {},
	}
}

func (btn *Button) Init() {
	btn.sprite = btn.parent.GetComponent(btn.sprite).(*Sprite)
}

func (btn *Button) Update() {
	if btn.parent.scene.Window.GetMouseDown() {
		fmt.Println(btn.parent.scene.Window.window.GetMouseButton(glfw.MouseButton1), prevMouseInput[glfw.MouseButton1])
		xi, yi := btn.parent.scene.Window.window.GetCursorPos()
		x := float32(xi)
		y := float32(yi)
		x1 := btn.parent.Position.X()
		y1 := -btn.parent.Position.Y()
		x2 := x1 + float32(btn.sprite.Texture.width)
		y2 := y1 + float32(btn.sprite.Texture.height)
		if x > x1 && y > y1 && x < x2 && y < y2 {
			fmt.Println(x, y, x1, y1, x2, y2)
			btn.OnClick()
		}
	}
}

func (btn *Button) SetParent(obj *Object) {
	btn.parent = obj
}
