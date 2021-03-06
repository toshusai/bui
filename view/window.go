package view

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

// Window !
type Window struct {
	window *glfw.Window
	Width  int
	Height int
	Update func()
	Scene  *Scene
}

// NewWindow create a new window
func NewWindow(width, height int, title string) *Window {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfwWindow, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	glfwWindow.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	window := &Window{
		window: glfwWindow,
		Update: func() {},
		Width:  width,
		Height: height,
	}
	return window
}

// AddScene add scene
func (w *Window) AddScene(scene *Scene) {
	w.Scene = scene
	scene.Window = w
}

// Run window loop
func (w *Window) Run() {
	defer glfw.Terminate()
	gl.Enable(gl.BLEND)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	prevMouseInput = map[glfw.MouseButton]glfw.Action{}
	for !w.window.ShouldClose() {
		w.Update()
		prevMouseInput[glfw.MouseButton1] = w.window.GetMouseButton(glfw.MouseButton1)
		w.window.SwapBuffers()
		glfw.PollEvents()
	}
}

// Close window
func (w *Window) Close() {
	w.window.SetShouldClose(true)
}

func (w *Window) GetMouseDown() bool {
	if prevMouseInput[glfw.MouseButton1] == glfw.Release &&
		w.window.GetMouseButton(glfw.MouseButton1) == glfw.Press {
		return true
	}
	return false
}

func (w *Window) GetMouseUp() bool {
	if prevMouseInput[glfw.MouseButton1] == glfw.Press &&
		w.window.GetMouseButton(glfw.MouseButton1) == glfw.Release {
		return true
	}
	return false
}

func (w *Window) GetCursorPos() (float64, float64) {
	return w.window.GetCursorPos()
}

func (w *Window) SetChar() {
	var c glfw.CharModsCallback
	c = func(w *glfw.Window, char rune, mod glfw.ModifierKey) {
		fmt.Println(string(char))
	}
	w.window.SetCharModsCallback(c)
}
