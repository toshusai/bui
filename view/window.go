package view

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()
}

// Window !
type Window struct {
	window *glfw.Window
	Update func()
}

// Run window loop
func (w *Window) Run() {
	defer glfw.Terminate()
	for !w.window.ShouldClose() {
		w.Update()

		w.window.SwapBuffers()
		glfw.PollEvents()
	}
}

// Close window
func (w *Window) Close() {
	w.window.SetShouldClose(true)
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
	}
	return window
}
