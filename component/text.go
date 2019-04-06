package component

import "github.com/toshusai/bui/view"

type Text struct {
	parent *view.Object
	Text   string
}

func (t *Text) Init() {}

func (t *Text) Update() {

}

func (t *Text) SetParent(obj *view.Object) {
	t.parent = obj
}
