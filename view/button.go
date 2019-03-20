package view

type Button struct {
	OnClick func()
	parent  *Object
}

func NewButton() Button {
	return Button{
		OnClick: func() {},
	}
}

func (btn *Button) Update() {
	// x, y := btn.parent.scene.Window.window.GetCursorPos()
}

func (btn *Button) SetParent(obj *Object) {
	btn.parent = obj
}
