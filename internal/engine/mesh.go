package engine

import "github.com/go-gl/mathgl/mgl32"	


type Mesh interface {
	Render(shader *Shader)
	Delete()
	GetTransform() *Transform
}

type BaseMesh struct {
	transform Transform
}

func NewBaseMesh() BaseMesh {
	return BaseMesh{
		transform: NewTransform(),
	}
}

func (b *BaseMesh) GetTransform() *Transform {
	return &b.transform
}

func (m *BaseMesh) SetPosition(x, y, z float32) {
	m.transform.Position = mgl32.Vec3{x, y, z}
}

func (m *BaseMesh) SetScale(x, y, z float32) {
	m.transform.Scale = mgl32.Vec3{x, y, z}
}

func (m *BaseMesh) SetRotation(angle float32, x, y, z float32) {
	m.transform.RotationAngle = angle
	m.transform.RotationAxis =  mgl32.Vec3{x, y, z}.Normalize()
}


