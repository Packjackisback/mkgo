package engine

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"	
)

type Engine struct {
	renderer *Renderer
	running bool
}

func New() *Engine {
	return &Engine {
		running: false,
	}
}

func (e *Engine) Run() error {
	fmt.Println("Starting Mario Kart Go")

	//creating the renderer
	var err error
  e.renderer, err = NewRenderer(800, 600, "Mario Kart Go")
  if err != nil {
    return fmt.Errorf("failed to create renderer: %v", err)
  }
  defer e.renderer.Cleanup()
	
	rect1 := NewRectangle()
	rect1.SetPosition(0.5, 0.0, 0.0)

	rect2 := NewRectangle()
	rect2.SetPosition(-0.5, 0.0, 0.0)
	rect2.SetScale(0.5, 0.5, 1.0)
	rect2.SetRotation(mgl32.DegToRad(45), 0, 0, 1) 

	e.renderer.AddMesh(rect1)
	e.renderer.AddMesh(rect2)


	e.running = true

	for !e.renderer.ShouldClose() && e.running {
        e.update()
        e.render()
  }

	return nil
}


func (e *Engine) update() {
    // Game logic will go here
}

func (e *Engine) render() {
    e.renderer.Clear()
    // Rendering calls will go here
    e.renderer.Render()

    e.renderer.SwapBuffers()
}
