package engine

import (
	"fmt"
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

	var err error
  e.renderer, err = NewRenderer(800, 600, "Mario Kart Go")
  if err != nil {
    return fmt.Errorf("failed to create renderer: %v", err)
  }
  defer e.renderer.Cleanup()
	
	
	circle1 := NewCircle(
		WithColor(0.99607843137, 0.890196078, 0.83137254),
		WithPosition(-0.2, -0.5, 0.0),
	)

	circle2 := NewCircle(
		WithColor(0.99607843137, 0.890196078, 0.83137254),
		WithPosition(0.2, -0.5, 0.0),
	)

	rect1 := NewRectangle(
		WithColor(0.99607843137, 0.890196078, 0.83137254),
		WithPosition(0.0, 0.0, 0.0),
		WithScale(0.7, 0.9, 0.0),
	)
	
	e.renderer.AddMesh(circle1)
	e.renderer.AddMesh(circle2)
	e.renderer.AddMesh(rect1)

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
