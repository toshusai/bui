package view

type Object struct {
	components []component
	scene      *Scene
}

func NewObject() Object {
	return Object{}
}

func (obj *Object) AddComponent(comp component) {
	comp.SetParent(obj)
	obj.components = append(obj.components, comp)
}
