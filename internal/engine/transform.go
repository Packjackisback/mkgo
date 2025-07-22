package engine

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	Position mgl32.Vec3
	RotationAxis mgl32.Vec3
	RotationAngle float32     // in degrees
	Scale    mgl32.Vec3
}

func NewTransform() Transform {
	return Transform{
		Position:      mgl32.Vec3{0, 0, 0},
		Scale:         mgl32.Vec3{1, 1, 1},
		RotationAxis:  mgl32.Vec3{0, 0, 1}, 
		RotationAngle: 0,
	}
}



func (t *Transform) GetModelMatrix() mgl32.Mat4 {
	model := mgl32.Ident4()

	model = model.Mul4(mgl32.Translate3D(t.Position.X(), t.Position.Y(), t.Position.Z()))
	model = model.Mul4(mgl32.HomogRotate3D(t.RotationAngle, t.RotationAxis))
	model = model.Mul4(mgl32.Scale3D(t.Scale.X(), t.Scale.Y(), t.Scale.Z()))

	return model
}


