package engine

type Rectangle struct {
	BaseMesh
}


func NewRectangle(opts ...MeshOption) *Rectangle {
	r := &Rectangle{BaseMesh: NewBaseMesh()}

	vertices := []float32{
		// Positions       // Colors
		 0.5,  0.5, 0.0,   1.0, 0.0, 0.0, // top right
		 0.5, -0.5, 0.0,   0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0,   0.0, 0.0, 1.0, // bottom left
		-0.5,  0.5, 0.0,   1.0, 1.0, 0.0, // top left
	}

	indices := []uint32{
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}

	r.SetVertexData(vertices, indices)

	for _, opt := range opts {
		opt(&r.BaseMesh)
	}

	return r
}

