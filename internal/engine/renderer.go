package engine

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Renderer struct {
	window *glfw.Window
	width  int
	height int
	shader *Shader
	meshes []Mesh
}

func NewRenderer(width, height int, title string) (*Renderer, error) {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize GLFW: %v", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %v", err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize OpenGL: %v", err)
	}

	gl.Viewport(0, 0, int32(width), int32(height))

	// create shader
	shader, err := NewShader()
	if err != nil {
		return nil, fmt.Errorf("failed to create shader: %v", err)
	}
	
	renderer := &Renderer{
		window: window,
		width:  width,
		height: height,
		shader: shader,
		meshes: make([]Mesh, 0),
	}

	

	return renderer, nil
}

func (r *Renderer) ShouldClose() bool {
	return r.window.ShouldClose()
}

func (r *Renderer) AddMesh(m Mesh) {
	r.meshes = append(r.meshes, m)
}

func (r *Renderer) SwapBuffers() {
	r.window.SwapBuffers()
	glfw.PollEvents()
}

func (r *Renderer) Clear() {
	gl.ClearColor(0.1, 0.1, 0.1, 1.0) // Dark gray background
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}


func (r *Renderer) Render() {
	r.shader.Use()

	for _, mesh := range r.meshes {
		model := mesh.GetTransform().ModelMatrix()
		r.shader.SetMat4("model", model)
		mesh.Render(r.shader)
	}
}

func (r *Renderer) Cleanup() {
	for _, mesh := range r.meshes {
		mesh.Delete()
	}
	r.shader.Delete()
	r.window.Destroy()
	glfw.Terminate()
}
