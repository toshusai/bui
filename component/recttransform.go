package component

import (
	"github.com/toshusai/bui/view"
)

type RectTransform struct {
	Width  float32
	Height float32
	parent *view.Object
}

func (rt *RectTransform) SetParent(obj *view.Object) {
	rt.parent = obj
}

func (rt *RectTransform) Update() {
}

func (rt *RectTransform) Init() {

}
