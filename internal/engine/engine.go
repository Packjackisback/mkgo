package engine

import (
	"fmt"
	"math"
	"github.com/go-gl/mathgl/mgl32"
	"time"
)

type Engine struct {
	renderer *Renderer
	camera 	 *Camera
	running bool
	startTime time.Time
}

func New() *Engine {
	return &Engine {
		running: false, camera: NewCamera(),
	}
}

func (e *Engine) Run() error {
	fmt.Println("Starting Mario Kart Go")

	var err error
  e.renderer, err = NewRenderer(800, 600, "Mario Kart Go", e.camera)
  if err != nil {
    return fmt.Errorf("failed to create renderer: %v", err)
  }
  defer e.renderer.Cleanup()
	
	
	
	model, err := LoadModel("internal/assets/Mario.obj", 
    WithPosition(0, 0, 0),
    WithScale(0.5, 0.5, 0.5),
    WithColor(1, 1, 1),
	)
	if err != nil {
    return fmt.Errorf("failed to load model: %v", err)
	}
	e.renderer.AddMesh(model)


	e.running = true
	e.startTime = time.Now()

	for !e.renderer.ShouldClose() && e.running {
        e.update()
        e.render()
  }

	return nil
}






func (e *Engine) update() {
	elapsed := time.Since(e.startTime).Seconds()
	
	if elapsed > 3.0 {
		fmt.Println("Rotating")
		radius := float32(5.0)
		speed := float32(0.5) 
		angle := float32(elapsed-3.0) * speed
		
		x := radius * float32(math.Cos(float64(angle)))
		z := radius * float32(math.Sin(float64(angle)))
		y := float32(2.0) 		
		e.camera.Position = mgl32.Vec3{x, y, z}
		
		e.camera.Target = mgl32.Vec3{0, 0, 0}
	}
}

func (e *Engine) render() {
    e.renderer.Clear()
    e.renderer.Render()
    e.renderer.SwapBuffers()
}
