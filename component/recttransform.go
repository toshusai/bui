package component

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/toshusai/bui"
	"github.com/toshusai/bui/view"
)

type RectTransformMode int

const (
	Center RectTransformMode = iota
)

type Anchors struct {
	Min bui.Vector2D
	Max bui.Vector2D
}

type RectTransform struct {
	Position            mgl32.Vec3
	Width               float32
	Height              float32
	parent              *view.Object
	Anchos              Anchors
	slippageX           float32
	slippageY           float32
	ParentRectTransform *RectTransform
}

func NewRectTransform() *RectTransform {
	return &RectTransform{
		Anchos: Anchors{
			Min: bui.Vector2D{},
			Max: bui.Vector2D{},
		},
	}
}

func (rt *RectTransform) GetAnchorsMinX() float32 {
	return rt.ParentRectTransform.Width * rt.Anchos.Min.X
}

func (rt *RectTransform) GetAnchorsMaxX() float32 {
	return rt.ParentRectTransform.Width * rt.Anchos.Max.X
}

func (rt *RectTransform) GetAnchorsMinY() float32 {
	return rt.ParentRectTransform.Height * rt.Anchos.Min.Y
}

func (rt *RectTransform) GetAnchorsMaxY() float32 {
	return rt.ParentRectTransform.Height * rt.Anchos.Max.Y
}

func (rt *RectTransform) IsPivotX() bool {
	return rt.Anchos.Min.X == rt.Anchos.Max.X
}

func (rt *RectTransform) IsPivotY() bool {
	return rt.Anchos.Min.Y == rt.Anchos.Max.Y
}

func (rt *RectTransform) SetParent(obj *view.Object) {
	rt.parent = obj
}

func (rt *RectTransform) Update() {

}

func (rt *RectTransform) Init() {
	if rt.parent.Parent != nil {
		rt.ParentRectTransform = rt.parent.Parent.GetComponent(&RectTransform{}).(*RectTransform)
	}
}
