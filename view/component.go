package view

type component interface {
	SetParent(*Object)
	Update()
	Init()
}
