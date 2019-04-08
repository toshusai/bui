package component

import (
	"fmt"

	"github.com/toshusai/bui/view"
)

type Button struct {
	OnPointerDwon func()
	OnPointerUp   func()
	OnDrag        func()
	parent        *view.Object
	rectTransform *RectTransform
	isPress       bool
}

func NewButton() *Button {
	f := func() {}
	return &Button{
		OnPointerDwon: f,
		OnPointerUp:   f,
		OnDrag:        f,
	}
}

func (btn *Button) Init() {
	btn.rectTransform = btn.parent.GetComponent(btn.rectTransform).(*RectTransform)
}

func (btn *Button) Update() {
	scene := view.GetScene()
	if scene.Window.GetMouseDown() || scene.Window.GetMouseUp() {
		xi, yi := scene.Window.GetCursorPos()
		x := float32(xi)
		y := float32(yi)
		x1 := btn.parent.Position.X()
		y1 := btn.parent.Position.Y()
		x2 := x1 + float32(btn.rectTransform.Width)
		y2 := y1 + float32(btn.rectTransform.Height)
		fmt.Println(x1, x2, y1, y2)
		if x > x1 && y > y1 && x < x2 && y < y2 {
			if scene.Window.GetMouseDown() {
				btn.OnPointerDwon()
				btn.isPress = true
			} else if btn.isPress && scene.Window.GetMouseUp() {
				btn.OnPointerUp()
				btn.isPress = false
			}
		}
	}
	if scene.Window.GetMouseUp() {
		btn.isPress = false
	}
	if btn.isPress {
		btn.OnDrag()
	}
}

func (btn *Button) SetParent(obj *view.Object) {
	btn.parent = obj
}
