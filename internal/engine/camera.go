package engine

import "github.com/go-gl/mathgl/mgl32"

type Camera struct {
	Position mgl32.Vec3
	Target   mgl32.Vec3
	Up       mgl32.Vec3
	FOV      float32
	Near     float32
	Far      float32
}

func NewCamera() *Camera {
	return &Camera{
		Position: mgl32.Vec3{0, 0, 3},
		Target:   mgl32.Vec3{0, 0, 0},
		Up:       mgl32.Vec3{0, 1, 0},
		FOV:      mgl32.DegToRad(45),
		Near:     0.1,
		Far:      100.0,
	}
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position, c.Target, c.Up)
}

func (c *Camera) GetProjectionMatrix(aspect float32) mgl32.Mat4 {
	return mgl32.Perspective(c.FOV, aspect, c.Near, c.Far)
}



