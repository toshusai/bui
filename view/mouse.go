package view

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var prevInput map[glfw.Key]glfw.Action
var prevMouseInput map[glfw.MouseButton]glfw.Action

func Update(w *Window) {
	prevInput[glfw.Key0] = w.window.GetKey(glfw.Key0)
}
