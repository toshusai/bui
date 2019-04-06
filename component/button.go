package component

import (
	"github.com/toshusai/bui/view"
)

type Button struct {
	OnClick func()
	parent  *view.Object
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
	scene := view.GetScene()
	if scene.Window.GetMouseDown() {
		xi, yi := scene.Window.GetCursorPos()
		x := float32(xi)
		y := float32(yi)
		x1 := btn.parent.Position.X()
		y1 := -btn.parent.Position.Y()
		x2 := x1 + float32(btn.sprite.Texture.GetWidth())
		y2 := y1 + float32(btn.sprite.Texture.GetHeight())
		if x > x1 && y > y1 && x < x2 && y < y2 {
			btn.OnClick()
		}
	}
}

func (btn *Button) SetParent(obj *view.Object) {
	btn.parent = obj
}
