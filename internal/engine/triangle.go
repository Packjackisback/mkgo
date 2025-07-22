package engine

type Triangle struct {
	BaseMesh
}


func NewTriangle(opts ...MeshOption) *Triangle {
	t := &Triangle{BaseMesh: NewBaseMesh()}

	vertices := []float32{
		0.0,  0.5, 0.0,   1.0, 0.0, 0.0,
		-0.5, -0.5, 0.0,   0.0, 1.0, 0.0,
		0.5, -0.5, 0.0,   0.0, 0.0, 1.0,
	}
	t.SetVertexData(vertices, nil)

	for _, opt := range opts {
		opt(&t.BaseMesh)
	}

	return t
}


